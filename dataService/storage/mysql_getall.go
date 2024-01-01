package storage

import (
	"context"
	"data-service/types"
	"database/sql"
	"fmt"
	"net/http"
)

func (s *MySqlStorage) RespondGetAll(ctx context.Context, rows *sql.Rows) (*types.ApiResponse, *types.ApiError) {
	entity := ctx.Value("entity").(string)
	switch entity {
	case "incidents":
		return s.respondGetAllIncidents(ctx, rows)
	case "incidenttypes":
		return s.respondGetAllIncidentTypes(ctx, rows)
	default:
		return nil, types.InternalServerError(fmt.Errorf("entity: %s not implemented", entity), ctx.Value("uri").(string))
	}
}

func (s *MySqlStorage) respondGetAllIncidents(ctx context.Context, rows *sql.Rows) (*types.ApiResponse, *types.ApiError) {
	defer rows.Close()
	incs := []types.Incident{}
	for rows.Next() {
		var inc types.Incident
		if err := inc.ScanTo(rows.Scan); err != nil {
			return nil, types.InternalServerError(err, ctx.Value("uri").(string))
		}
		incs = append(incs, inc)
	}
	return types.NewApiResponse(http.StatusOK, ctx.Value("uri").(string), incs), nil
}

func (s *MySqlStorage) respondGetAllIncidentTypes(ctx context.Context, rows *sql.Rows) (*types.ApiResponse, *types.ApiError) {
	defer rows.Close()
	iTs := []types.IncidentType{}
	for rows.Next() {
		var iT types.IncidentType
		if err := iT.ScanTo(rows.Scan); err != nil {
			return nil, types.InternalServerError(err, ctx.Value("uri").(string))
		}
		iTs = append(iTs, iT)
	}
	return types.NewApiResponse(http.StatusOK, ctx.Value("uri").(string), iTs), nil
}
