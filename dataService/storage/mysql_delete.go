package storage

import (
	"context"
	"data-service/types"
	"fmt"
	"net/http"
)

func (s *MySqlStorage) Delete(ctx context.Context, entity, id string) (*types.ApiResponse, *types.ApiError) {
	switch entity {
	case "incidents":
		return s.deleteIncident(ctx, id)
	case "incidenttypes":
		return s.deleteIncidentType(ctx, id)
	default:
		return nil, types.InternalServerError(fmt.Errorf("entity: %s not implemented", entity), ctx.Value("uri").(string))
	}
}

func (s *MySqlStorage) deleteIncident(ctx context.Context, id string) (*types.ApiResponse, *types.ApiError) {
	// TODO implemet
	return nil, types.NewApiError(http.StatusNotImplemented, ctx.Value("uri").(string), fmt.Errorf("delete incident: %s not implemented", id))
}

func (s *MySqlStorage) deleteIncidentType(ctx context.Context, id string) (*types.ApiResponse, *types.ApiError) {
	// TODO implemet
	return nil, types.NewApiError(http.StatusNotImplemented, ctx.Value("uri").(string), fmt.Errorf("delete incident: %s not implemented", id))
}
