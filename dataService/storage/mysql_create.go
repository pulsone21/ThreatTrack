package storage

import (
	"context"
	"data-service/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (s *MySqlStorage) CreateEntity(ctx context.Context, sql string, body io.ReadCloser) (*types.ApiResponse, *types.ApiError) {
	entity := ctx.Value("entity").(string)
	switch entity {
	case "incident":
		return s.createIncident(ctx, body, sql)
	case "incidenttype":
		return s.createIncidentType(ctx, body, sql)
	default:
		return nil, types.InternalServerError(fmt.Errorf("entity: %s not implemented", entity), ctx.Value("uri").(string))
	}
}

func (s *MySqlStorage) createIncident(ctx context.Context, body io.ReadCloser, sql string) (*types.ApiResponse, *types.ApiError) {
	var inc types.Incident
	err := json.NewDecoder(body).Decode(&inc)
	if err != nil {
		return nil, types.BadRequestError(fmt.Errorf("bad payload for request\n%s", err.Error()), ctx.Value("uri").(string))
	}
	_, err = s.Db.Exec(sql, inc.Id, inc.Name, inc.Severity, inc.IncidentType.Id)
	if err != nil {
		return nil, types.InternalServerError(err, ctx.Value("uri").(string))
	}
	return types.NewApiResponse(http.StatusCreated, ctx.Value("uri").(string), inc), nil
}

func (s *MySqlStorage) createIncidentType(ctx context.Context, body io.ReadCloser, sql string) (*types.ApiResponse, *types.ApiError) {
	var inc types.IncidentType
	err := json.NewDecoder(body).Decode(&inc)
	if err != nil {
		return nil, types.BadRequestError(fmt.Errorf("bad payload for request\n%s", err.Error()), ctx.Value("uri").(string))
	}
	_, err = s.Db.Exec(sql, inc.Name)
	if err != nil {
		return nil, types.InternalServerError(err, ctx.Value("uri").(string))
	}
	return types.NewApiResponse(http.StatusCreated, ctx.Value("uri").(string), inc), nil
}
