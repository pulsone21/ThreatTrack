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
	EntityStore[*types.IncidentType]
	db *sql.DB
}

func NewIncidentTypeStore(storage *MySqlStorage) *IncidentTypeStore {
	return &IncidentTypeStore{
		storage: storage,
		db:      storage.Db,
	}
}

func (i *IncidentTypeStore) Get(ctx context.Context, id string) (*types.IncidentType, *types.ApiError) {
	uri := ctx.Value("uri").(string)
	loadedSql, err := LoadRawSQL("incidenttypes/GetById.sql")
	if err != nil {
		return nil, types.InternalServerError(err, uri)
	}
	res := i.db.QueryRow(loadedSql, id)
	if res.Err() != nil {
		if res.Err() == sql.ErrNoRows {
			return nil, types.NotFoundError(fmt.Errorf("no incident found"), uri)
		}
		return nil, types.InternalServerError(res.Err(), uri)
	}
	var iT types.IncidentType
	if err := iT.ScanTo(res.Scan); err != nil {
		return nil, types.InternalServerError(err, ctx.Value("uri").(string))
	}
	return &iT, nil
}

func (i *IncidentTypeStore) GetAll(ctx context.Context, qP QueryParameter) (*[]types.IncidentType, *types.ApiError) {
	uri := ctx.Value("uri").(string)
	loadedSql, err := LoadRawSQL("incidenttypes/GetAll.sql")
	if err != nil {
		return nil, types.InternalServerError(err, uri)
	}
	res, err := i.db.Query(loadedSql, qP.Limit, qP.Offset)
	if err != nil {
		return nil, types.InternalServerError(err, uri)
	}
	if res.Err() != nil {
		if res.Err() == sql.ErrNoRows {
			return nil, types.NotFoundError(fmt.Errorf("no incidents found"), uri)
		}
		return nil, types.InternalServerError(res.Err(), uri)
	}
	defer res.Close()
	iTs := []types.IncidentType{}
	for res.Next() {
		var iT types.IncidentType
		if err := iT.ScanTo(res.Scan); err != nil {
			return nil, types.InternalServerError(err, uri)
		} else {
			iTs = append(iTs, iT)
		}
	}
	return &iTs, nil
}

func (i *IncidentTypeStore) GetQuery(ctx context.Context, qP QueryParameter) (*[]types.IncidentType, *types.ApiError) {
	// ! This Entity isn't queryable
	return nil, types.BadRequestError(fmt.Errorf("not applicable"), "/incidenttypes/query")
}

func (i *IncidentTypeStore) Create(ctx context.Context, incidentType types.IncidentType) (*types.IncidentType, *types.ApiError) {
	fmt.Println("creating new inc from ", incidentType)
	uri := ctx.Value("uri").(string)
	loadedSql, err := LoadRawSQL("incidenttypes/Create.sql")
	if err != nil {
		return nil, types.InternalServerError(err, uri)
	}
	res, err := i.db.Exec(loadedSql, incidentType.Name)
	if err != nil {
		return nil, types.InternalServerError(err, uri)
	}
	//The Id is auto incrementing in the database, so the given id from the parameters are irrelevant,
	//since we now the acutal id only after inserting to db.
	iD, _ := res.LastInsertId()
	iT := types.IncidentType{
		Name: incidentType.Name,
		Id:   iD,
	}
	return &iT, nil
}

func (i *IncidentTypeStore) Update(ctx context.Context, entity types.IncidentType) (*types.IncidentType, *types.ApiError) {
	uri := ctx.Value("uri").(string)
	return nil, types.NotImplementedError(fmt.Errorf("not implemented"), uri)
}

func (i *IncidentTypeStore) Delete(ctx context.Context, id string) *types.ApiError {
	uri := ctx.Value("uri").(string)
	loadedSql, err := LoadRawSQL("incidenttypes/Delete.sql")
	if err != nil {
		return types.InternalServerError(err, uri)
	}
	_, err = i.db.Exec(loadedSql, id)
	if err != nil {
		return types.InternalServerError(err, uri)
	}
	return nil
}
