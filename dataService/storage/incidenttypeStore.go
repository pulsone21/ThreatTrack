package storage

import (
	"context"
	"data-service/types"
	"database/sql"
	"fmt"
)

type RequestIncidentType struct {
	Name string
}
type IncidentTypeStore struct {
	storage *MySqlStorage
	EntityStoreInterface[*types.IncidentType]
	db *sql.DB
}

func NewIncidentTypeStore(storage *MySqlStorage) *IncidentTypeStore {
	return &IncidentTypeStore{
		storage: storage,
		db:      storage.Db,
	}
}

func (iT *IncidentTypeStore) Get(ctx context.Context, id string) (*types.IncidentType, *types.ApiError) {
	panic("not implemented") // TODO: Implement
}

func (iT *IncidentTypeStore) GetAll(ctx context.Context, qP QueryParameter) (*[]types.IncidentType, *types.ApiError) {
	panic("not implemented") // TODO: Implement
}

func (iT *IncidentTypeStore) GetQuery(ctx context.Context, qP QueryParameter) (*[]types.IncidentType, *types.ApiError) {
	return nil, types.BadRequestError(fmt.Errorf("not applicable"), "/incidenttypes/query")
}

func (iT *IncidentTypeStore) Create(ctx context.Context, incidentType types.IncidentType) (*types.IncidentType, *types.ApiError) {
	panic("not implemented") // TODO: Implement
}

func (iT *IncidentTypeStore) Update(ctx context.Context, entity types.IncidentType) (*types.IncidentType, *types.ApiError) {
	panic("not implemented") // TODO: Implement
}

func (iT *IncidentTypeStore) Delete(ctx context.Context, id string) *types.ApiError {
	panic("not implemented") // TODO: Implement
}
