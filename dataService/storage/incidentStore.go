package storage

import (
	"context"
	"data-service/types"
	"database/sql"
	"fmt"
)

type RequestIncident struct {
	Name         string
	Severity     string
	IncidentType int
}

type IncidentStore struct {
	storage *MySqlStorage
	EntityStore[*types.Incident]
	db *sql.DB
}

func NewIncidentStore(storage *MySqlStorage) *IncidentStore {
	return &IncidentStore{
		storage: storage,
		db:      storage.Db,
	}
}

func (i *IncidentStore) Get(ctx context.Context, id string) (*types.Incident, *types.ApiError) {
	uri := ctx.Value("uri").(string)
	loadedSql, err := LoadRawSQL("incidents/GetById.sql")
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
	var inc types.Incident
	if err := inc.ScanTo(res.Scan); err != nil {
		return nil, types.InternalServerError(err, uri)
	}
	return &inc, nil
}

func (i *IncidentStore) GetAll(ctx context.Context, qP QueryParameter) (*[]types.Incident, *types.ApiError) {
	uri := ctx.Value("uri").(string)
	loadedSql, err := LoadRawSQL("incidents/GetAll.sql")
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
	incs := []types.Incident{}
	for res.Next() {
		var inc types.Incident
		if err := inc.ScanTo(res.Scan); err != nil {
			return nil, types.InternalServerError(err, uri)
		} else {
			incs = append(incs, inc)
		}
	}
	return &incs, nil
}

func (i *IncidentStore) GetQuery(ctx context.Context, qP QueryParameter) (*[]types.Incident, *types.ApiError) {
	uri := ctx.Value("uri").(string)
	rawSql, err := LoadRawSQL("incidents/GetQuery.sql")
	if err != nil {
		return nil, types.InternalServerError(err, uri)
	}
	whiteList := i.createWhitelist()
	if whiteList == nil {
		return nil, types.InternalServerError(fmt.Errorf("couldn't create whitelist for entity"), uri)
	}
	for key, val := range qP.Query {
		if !CheckWhitelist(key, val, whiteList) {
			return nil, types.BadRequestError(fmt.Errorf("whitelist check failed"), uri)
		}
	}
	finalSql := FinalizeSQL(rawSql, "incidents", qP)
	rows, err := i.db.Query(finalSql)
	if err != nil {
		return nil, types.InternalServerError(err, uri)
	}
	if rows.Err() != nil {
		if rows.Err() == sql.ErrNoRows {
			return nil, types.NotFoundError(fmt.Errorf("no incidents found"), uri)
		}
		return nil, types.InternalServerError(rows.Err(), uri)
	}
	defer rows.Close()
	var incs []types.Incident
	for rows.Next() {
		var i types.Incident
		err := i.ScanTo(rows.Scan)
		if err != nil {
			return nil, types.InternalServerError(rows.Err(), uri)
		}
		incs = append(incs, i)
	}
	return &incs, nil
}

func (i *IncidentStore) Create(ctx context.Context, inc *types.Incident) (*types.Incident, *types.ApiError) {
	fmt.Println("creating new inc from ", inc)
	uri := ctx.Value("uri").(string)
	loadedSql, err := LoadRawSQL("incidents/Create.sql")
	if err != nil {
		return nil, types.InternalServerError(err, uri)
	}
	_, err = i.db.Exec(loadedSql, inc.Id, inc.Name, inc.Severity, inc.IncidentType.Id)
	if err != nil {
		return nil, types.InternalServerError(err, uri)
	}
	return inc, nil
}

func (i *IncidentStore) Update(entity types.Incident) (*types.Incident, *types.ApiError) {
	panic("not implemented") // TODO: Implement
}

func (i *IncidentStore) Delete(ctx context.Context, id string) *types.ApiError {
	uri := ctx.Value("uri").(string)
	loadedSql, err := LoadRawSQL("incidents/Delete.sql")
	if err != nil {
		return types.InternalServerError(err, uri)
	}
	_, err = i.db.Exec(loadedSql, id)
	if err != nil {
		return types.InternalServerError(err, uri)
	}
	return nil
}

func (i *IncidentStore) createWhitelist() Whitelist {
	sql, err := LoadRawSQL("incidenttypes/GetAll.sql")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := i.db.Query(sql)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Close()
	iTs := []string{}

	for res.Next() {
		var iT types.IncidentType
		err = iT.ScanTo(res.Scan)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		iTs = append(iTs, iT.Name)
	}
	incWhitelist := map[string][]string{
		"Severity": {string(types.Low), string(types.Medium), string(types.High), string(types.Critical)},
		"Status":   {string(types.Pending), string(types.Open), string(types.Active), string(types.Closed)},
		"Type":     iTs,
	}
	return incWhitelist
}
