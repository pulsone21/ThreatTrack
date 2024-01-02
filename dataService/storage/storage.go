package storage

import (
	"context"
	"data-service/types"
	"net/http"
)

type DBConfig struct {
	Address  string
	Port     string
	User     string
	Password string
	Database string
}

type Storage interface {
	HandleGetAll(context.Context, http.ResponseWriter, *http.Request) (*types.ApiResponse, *types.ApiError)
	HandleGetID(context.Context, http.ResponseWriter, *http.Request) (*types.ApiResponse, *types.ApiError)
	HandleGetQuery(context.Context, http.ResponseWriter, *http.Request) (*types.ApiResponse, *types.ApiError)
	HandleCreate(context.Context, http.ResponseWriter, *http.Request) (*types.ApiResponse, *types.ApiError)
	HandleDelete(context.Context, http.ResponseWriter, *http.Request) (*types.ApiResponse, *types.ApiError)
	HandleUpdate(context.Context, http.ResponseWriter, *http.Request) (*types.ApiResponse, *types.ApiError)
}

type EntityStore[T types.Entity] interface {
	Get(ctx context.Context, id string) (*T, *types.ApiError)
	GetAll(ctx context.Context, qP QueryParameter) (*[]T, *types.ApiError)
	GetQuery(ctx context.Context, qP QueryParameter) (*[]T, *types.ApiError)
	Create(ctx context.Context, entity T) (*T, *types.ApiError)
	Update(ctx context.Context, entity T) (*T, *types.ApiError)
	Delete(ctx context.Context, id string) *types.ApiError
}

type Whitelist map[string][]string
