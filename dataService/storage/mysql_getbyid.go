package storage

import (
	"context"
	"data-service/types"
	"database/sql"
	"fmt"
	"net/http"
)

func (s *MySqlStorage) RespondGetId(ctx context.Context, row *sql.Row) (*types.ApiResponse, *types.ApiError) {
	entity := ctx.Value("entity").(string)
	switch entity {
	case "incidents":
		return s.respondGetIncidentById(ctx, row)
	case "incidenttypes":
		return s.respondGetIncidentTypeById(ctx, row)
	default:
		return nil, types.InternalServerError(fmt.Errorf("entity: %s not implemented", entity), ctx.Value("uri").(string))
	}
}

func (s *MySqlStorage) respondGetIncidentById(ctx context.Context, row *sql.Row) (*types.ApiResponse, *types.ApiError) {
	var inc types.Incident
	if err := inc.ScanTo(row.Scan); err != nil {
		return nil, types.InternalServerError(err, ctx.Value("uri").(string))
	}
	return types.NewApiResponse(http.StatusOK, ctx.Value("uri").(string), inc), nil
}
func (s *MySqlStorage) respondGetIncidentTypeById(ctx context.Context, rows *sql.Row) (*types.ApiResponse, *types.ApiError) {
	var iT types.IncidentType
	if err := iT.ScanTo(rows.Scan); err != nil {
		return nil, types.InternalServerError(err, ctx.Value("uri").(string))
	}
	return types.NewApiResponse(http.StatusOK, ctx.Value("uri").(string), iT), nil
}
