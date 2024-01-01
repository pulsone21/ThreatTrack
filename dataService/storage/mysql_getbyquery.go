package storage

import (
	"context"
	"data-service/types"
	"database/sql"
	"fmt"
	"net/http"
)

func (s *MySqlStorage) RespondQuery(ctx context.Context, w http.ResponseWriter, r *http.Request) (*types.ApiResponse, *types.ApiError) {
	entity := ctx.Value("entity").(string)
	uri := ctx.Value("uri").(string)
	sql, err := LoadRawSQL(fmt.Sprintf("%s/GetAll.sql", entity))
	if err != nil {
		return nil, types.InternalServerError(err, r.RequestURI)
	}
	withParams := true
	qP := NewQueryParameter(r.URL.Query(), withParams)
	for key, val := range qP.Query {
		if !s.CheckWhitelist(entity, key, val) {
			return nil, types.BadRequestError(fmt.Errorf("whitelist check failed"), r.RequestURI)
		}
	}
	finalSql := FinalizeSQL(sql, entity, *qP)
	rows, err := s.Db.Query(finalSql)
	if err != nil {
		return nil, types.InternalServerError(err, r.RequestURI)
	}
	if rows.Err() != nil {
		if rows.Err() == sql.ErrNoRows {
			return nil, types.NotFoundError(fmt.Errorf("no %s found", entity), r.RequestURI)
		}
		return nil, types.InternalServerError(rows.Err(), r.RequestURI)
	}
	switch entity {
	case "incidents":
		incs, err := s.scanToIncident(rows)
		if err != nil {
			return nil, types.InternalServerError(err, uri)
		}
		return types.NewApiResponse(http.StatusOK, uri, incs), nil
	default:
		return nil, types.BadRequestError(fmt.Errorf("entity: %s not queryable", entity), uri)
	}
}

func (s *MySqlStorage) scanToIncident(rows *sql.Rows) (*[]types.Incident, error) {
	var incs []types.Incident
	for rows.Next() {
		var i types.Incident
		err := i.ScanTo(rows.Scan)
		if err != nil {
			return nil, err
		}
		incs = append(incs, i)
	}
	return &incs, nil
}
