package storage

import (
	"context"
	"data-service/types"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type MySqlStorage struct {
	Db                *sql.DB
	IncidentStore     *IncidentStore
	IncidentTypeStore *IncidentTypeStore
	UserStore         *UserStore
	TaskStore         *TaskStore
}

func NewMySqlStorage(dbConf DBConfig) *MySqlStorage {
	storage := &MySqlStorage{}
	storage.setUpDB(dbConf)
	storage.IncidentStore = NewIncidentStore(storage)
	storage.IncidentTypeStore = NewIncidentTypeStore(storage)
	storage.UserStore = NewUserStore(storage)
	storage.TaskStore = NewTaskStore(storage)
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
	var value interface{}
	var err *types.ApiError
	withParams := false
	qP := NewQueryParameter(r.URL.Query(), withParams)
	switch entity {
	case "incidents":
		value, err = s.IncidentStore.GetAll(ctx, *qP)
	case "incidenttypes":
		value, err = s.IncidentTypeStore.GetAll(ctx, *qP)
	case "users":
		value, err = s.UserStore.GetAll(ctx, *qP)
	case "tasks":
		value, err = s.TaskStore.GetAll(ctx, *qP)
	default:
		return nil, types.NotImplementedError(fmt.Errorf("entity: %s not implemented", entity), ctx.Value("uri").(string))
	}
	if err != nil {
		return nil, err
	}
	return types.NewApiResponse(http.StatusOK, ctx.Value("uri").(string), value), nil
}

func (s *MySqlStorage) HandleGetQuery(ctx context.Context, w http.ResponseWriter, r *http.Request) (*types.ApiResponse, *types.ApiError) {
	entity := ctx.Value("entity").(string)
	uri := ctx.Value("uri").(string)
	var value interface{}
	var err *types.ApiError
	withParams := true
	qP := NewQueryParameter(r.URL.Query(), withParams)
	switch entity {
	case "incidents":
		value, err = s.IncidentStore.GetQuery(ctx, *qP)
	case "tasks":
		value, err = s.TaskStore.GetQuery(ctx, *qP)
	default:
		return nil, types.NotImplementedError(fmt.Errorf("entity: %s not implemented", entity), uri)
	}
	if err != nil {
		return nil, err
	}
	return types.NewApiResponse(http.StatusOK, ctx.Value("uri").(string), value), nil
}

func (s *MySqlStorage) HandleGetID(ctx context.Context, w http.ResponseWriter, r *http.Request) (*types.ApiResponse, *types.ApiError) {
	entity := ctx.Value("entity").(string)
	id := mux.Vars(r)["id"]
	var value interface{}
	var err *types.ApiError
	switch entity {
	case "incidents":
		value, err = s.IncidentStore.Get(ctx, id)
	case "incidenttypes":
		value, err = s.IncidentTypeStore.Get(ctx, id)
	case "users":
		value, err = s.UserStore.Get(ctx, id)
	case "tasks":
		value, err = s.TaskStore.Get(ctx, id)
	default:
		return nil, types.NotImplementedError(fmt.Errorf("not implemented"), ctx.Value("uri").(string))
	}
	if err != nil {
		return nil, err
	}
	return types.NewApiResponse(http.StatusOK, ctx.Value("uri").(string), value), nil
}

func (s *MySqlStorage) HandleCreate(ctx context.Context, w http.ResponseWriter, r *http.Request) (*types.ApiResponse, *types.ApiError) {
	entity := ctx.Value("entity").(string)
	uri := ctx.Value("uri").(string)
	var value interface{}
	var err *types.ApiError
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	switch entity {
	case "incidents":
		var iR RequestIncident
		dec.Decode(&iR)
		iT, err := s.IncidentTypeStore.Get(ctx, fmt.Sprint(iR.IncidentType))
		if err != nil {
			return nil, err
		}
		value = types.NewIncident(iR.Name, types.Priority(iR.Severity), *iT)
		_, err = s.IncidentStore.Create(ctx, value.(*types.Incident))
	case "incidenttypes":
		var iR RequestIncidentType
		dec.Decode(&iR)
		iT := types.NewIncidentType(iR.Name)
		value, err = s.IncidentTypeStore.Create(ctx, *iT)
	case "users":
		var uR RequestUser
		dec.Decode(&uR)
		usr := types.NewUser(uR.Firstname, uR.Lastname, uR.Email)
		value, err = s.UserStore.Create(ctx, *usr)
	case "tasks":
		var tR RequestTask
		dec.Decode(&tR)
		owner, err := s.UserStore.Get(ctx, tR.OwnerId)
		if err != nil {
			return nil, types.BadRequestError(fmt.Errorf("User with id %s not found.", tR.OwnerId), uri)
		}
		inc_id, err1 := uuid.Parse(tR.IncidentID)
		if err1 != nil {
			return nil, types.BadRequestError(fmt.Errorf("IncidentID %s not an UUID.", tR.IncidentID), uri)
		}
		value = types.NewTask(tR.Title, tR.Description, *owner, inc_id)
		_, err = s.TaskStore.Create(ctx, value.(*types.Task))
	default:
		return nil, types.NotImplementedError(fmt.Errorf("not implemented"), uri)
	}
	if err != nil {
		return nil, err
	}
	return types.NewApiResponse(http.StatusOK, uri, value), nil
}
func (s *MySqlStorage) HandleDelete(ctx context.Context, w http.ResponseWriter, r *http.Request) (*types.ApiResponse, *types.ApiError) {
	entity := ctx.Value("entity").(string)
	var err *types.ApiError
	uri := ctx.Value("uri").(string)
	id := mux.Vars(r)["id"]
	switch entity {
	case "incidents":
		err = s.IncidentStore.Delete(ctx, id)
	case "incidenttypes":
		err = s.IncidentTypeStore.Delete(ctx, id)
	case "users":
		err = s.UserStore.Delete(ctx, id)
	case "tasks":
		err = s.TaskStore.Delete(ctx, id)
	default:
		return nil, types.NotImplementedError(fmt.Errorf("not implemented"), uri)
	}
	if err != nil {
		return nil, err
	}
	return types.NewApiResponse(http.StatusOK, uri, fmt.Sprintf("%s with id: %s deleted", entity, id)), nil
}

func (s *MySqlStorage) HandleUpdate(ctx context.Context, w http.ResponseWriter, r *http.Request) (*types.ApiResponse, *types.ApiError) {
	uri := ctx.Value("uri").(string)
	// TODO Implement
	return nil, types.NotImplementedError(fmt.Errorf("not implemented"), uri)
}
