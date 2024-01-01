package storage

import (
	"context"
	"data-service/types"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlStorage struct {
	Db *sql.DB
}

func NewMySqlStorage(dbConf DBConfig) *MySqlStorage {
	storage := &MySqlStorage{}
	storage.setUpDB(dbConf)
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
	loadedSql, err := LoadRawSQL(fmt.Sprintf("%s/GetAll.sql", entity))
	if err != nil {
		return nil, types.InternalServerError(err, r.RequestURI)
	}
	withParams := false
	qP := NewQueryParameter(r.URL.Query(), withParams)
	res, err := s.Db.Query(loadedSql, qP.Limit, qP.Offset)
	if err != nil {
		return nil, types.InternalServerError(err, r.RequestURI)
	}
	if res.Err() != nil {
		if res.Err() == sql.ErrNoRows {
			return nil, types.NotFoundError(fmt.Errorf("no %s found", entity), r.RequestURI)
		}
		return nil, types.InternalServerError(res.Err(), r.RequestURI)
	}
	return s.RespondGetAll(ctx, res)
}

func (s *MySqlStorage) HandleGetQuery(ctx context.Context, w http.ResponseWriter, r *http.Request) (*types.ApiResponse, *types.ApiError) {
	return nil, types.NewApiError(http.StatusNotImplemented, r.RequestURI, fmt.Errorf("not Implemented"))
	return s.RespondQuery(ctx, w, r)
}

func (s *MySqlStorage) HandleGetID(ctx context.Context, w http.ResponseWriter, r *http.Request) (*types.ApiResponse, *types.ApiError) {
	entity := ctx.Value("entity").(string)
	id := r.URL.Path[len(r.URL.Path)-1]
	loadedSql, err := LoadRawSQL(fmt.Sprintf("%s/GetById.sql", entity))
	if err != nil {
		return nil, types.InternalServerError(err, r.RequestURI)
	}
	res := s.Db.QueryRow(loadedSql, id)
	if res.Err() != nil {
		if res.Err() == sql.ErrNoRows {
			return nil, types.NotFoundError(fmt.Errorf("no %s found", entity), r.RequestURI)
		}
		return nil, types.InternalServerError(res.Err(), r.RequestURI)
	}
	return s.RespondGetId(ctx, res)
}

func (s *MySqlStorage) HandleCreate(ctx context.Context, w http.ResponseWriter, r *http.Request) (*types.ApiResponse, *types.ApiError) {
	entity := ctx.Value("entity").(string)
	rawSql, err := LoadRawSQL(fmt.Sprintf("%s/create.sql", entity))
	if err != nil {
		return nil, types.InternalServerError(err, r.RequestURI)
	}
	return s.CreateEntity(ctx, rawSql, r.Body)
}
func (s *MySqlStorage) HandleDelete(ctx context.Context, w http.ResponseWriter, r *http.Request) (*types.ApiResponse, *types.ApiError) {
	return nil, nil
}
func (s *MySqlStorage) HandleUpdate(ctx context.Context, w http.ResponseWriter, r *http.Request) (*types.ApiResponse, *types.ApiError) {
	return nil, nil
}
