package storage

import (
	"context"
	"data-service/types"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type MySqlStorage struct {
	Db                *sql.DB
	IncidentStore     *IncidentStore
	IncidentTypeStore *IncidentTypeStore
}

func NewMySqlStorage(dbConf DBConfig) *MySqlStorage {
	storage := &MySqlStorage{}
	storage.setUpDB(dbConf)
	storage.IncidentStore = NewIncidentStore(storage)
	storage.IncidentTypeStore = NewIncidentTypeStore(storage)
	return storage
}

func (s *MySqlStorage) setUpDB(dbConf DBConfig) {

	fmt.Printf("Connecting to MySQL at %s:%s\n", dbConf.Address, dbConf.Port)
	db_addres := fmt.Sprintf("%s:%s@tcp(%s:%s)/", dbConf.User, dbConf.Password, dbConf.Address, dbConf.Port)

	db, err := sql.Open("mysql", db_addres)
	if err != nil {
		panic(err)
	}
	fmt.Println("Initial Contact made")
	creatDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbConf.Database)
	_, err = db.Exec(creatDB)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created Database %s\n", creatDB)
	connect := fmt.Sprintf("%s%s?parseTime=true", db_addres, dbConf.Database)
	db, err = sql.Open("mysql", connect)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to Database, with %s\n", connect)
	s.Db = db
}

func (s *MySqlStorage) HandleGetAll(ctx context.Context, w http.ResponseWriter, r *http.Request) (*types.ApiResponse, *types.ApiError) {
	entity := ctx.Value("entity").(string)
	withParams := false
	qP := NewQueryParameter(r.URL.Query(), withParams)
	switch entity {
	case "incidents":
		incs, err := s.IncidentStore.GetAll(ctx, *qP)
		if err != nil {
			return nil, err
		}
		return types.NewApiResponse(http.StatusOK, ctx.Value("uri").(string), incs), nil
	case "incidenttypes":
		iTs, err := s.IncidentTypeStore.GetAll(ctx, *qP)
		if err != nil {
			return nil, err
		}
		return types.NewApiResponse(http.StatusOK, ctx.Value("uri").(string), iTs), nil
	default:
		return nil, types.NotImplementedError(fmt.Errorf("entity: %s not implemented", entity), ctx.Value("uri").(string))
	}
}

func (s *MySqlStorage) HandleGetQuery(ctx context.Context, w http.ResponseWriter, r *http.Request) (*types.ApiResponse, *types.ApiError) {
	entity := ctx.Value("entity").(string)
	uri := ctx.Value("uri").(string)
	withParams := true
	qP := NewQueryParameter(r.URL.Query(), withParams)
	switch entity {
	case "incidents":
		incs, err := s.IncidentStore.GetQuery(ctx, *qP)
		if err != nil {
			return nil, err
		}
		return types.NewApiResponse(http.StatusOK, uri, incs), nil
	default:
		return nil, types.NotImplementedError(fmt.Errorf("entity: %s not implemented", entity), uri)
	}
}

func (s *MySqlStorage) HandleGetID(ctx context.Context, w http.ResponseWriter, r *http.Request) (*types.ApiResponse, *types.ApiError) {
	entity := ctx.Value("entity").(string)
	id := mux.Vars(r)["id"]
	switch entity {
	case "incidents":
		inc, err := s.IncidentStore.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		return types.NewApiResponse(http.StatusOK, ctx.Value("uri").(string), inc), nil
	case "incidenttypes":
		iT, err := s.IncidentTypeStore.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		return types.NewApiResponse(http.StatusOK, ctx.Value("uri").(string), iT), nil
	default:
		return nil, types.NotImplementedError(fmt.Errorf("not implemented"), ctx.Value("uri").(string))
	}
}

func (s *MySqlStorage) HandleCreate(ctx context.Context, w http.ResponseWriter, r *http.Request) (*types.ApiResponse, *types.ApiError) {
	entity := ctx.Value("entity").(string)
	uri := ctx.Value("uri").(string)
	switch entity {
	case "incidents":
		var iR RequestIncident
		json.NewDecoder(r.Body).Decode(&iR)
		iT, err := s.IncidentTypeStore.Get(ctx, fmt.Sprint(iR.IncidentType))
		if err != nil {
			return nil, err
		}
		inc := types.NewIncident(iR.Name, types.IncidentSeverity(iR.Severity), *iT)
		if _, err := s.IncidentStore.Create(ctx, inc); err != nil {
			return nil, err
		}
		return types.NewApiResponse(http.StatusOK, uri, inc), nil
	case "incidenttypes":
		var iR RequestIncidentType
		json.NewDecoder(r.Body).Decode(&iR)
		iT := types.NewIncidentType(iR.Name)
		if _, err := s.IncidentTypeStore.Create(ctx, *iT); err != nil {
			return nil, err
		}
		return types.NewApiResponse(http.StatusOK, uri, iT), nil
	default:
		return nil, types.NotImplementedError(fmt.Errorf("not implemented"), uri)
	}
}
func (s *MySqlStorage) HandleDelete(ctx context.Context, w http.ResponseWriter, r *http.Request) (*types.ApiResponse, *types.ApiError) {
	entity := ctx.Value("entity").(string)
	uri := ctx.Value("uri").(string)
	id := mux.Vars(r)["id"]
	switch entity {
	case "incidents":
		if err := s.IncidentStore.Delete(ctx, id); err != nil {
			return nil, err
		}
		return types.NewApiResponse(http.StatusOK, uri, fmt.Sprintf("Incident with id: %s deleted", id)), nil
	case "incidenttypes":
		if err := s.IncidentTypeStore.Delete(ctx, id); err != nil {
			return nil, err
		}
		return types.NewApiResponse(http.StatusOK, uri, fmt.Sprintf("IncidentType with id: %s deleted", id)), nil
	default:
		return nil, types.NotImplementedError(fmt.Errorf("not implemented"), uri)
	}
}
func (s *MySqlStorage) HandleUpdate(ctx context.Context, w http.ResponseWriter, r *http.Request) (*types.ApiResponse, *types.ApiError) {
	uri := ctx.Value("uri").(string)
	return nil, types.NotImplementedError(fmt.Errorf("not implemented"), uri)
}
