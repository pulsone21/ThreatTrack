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
	CheckWhitelist(string, string, string) bool
}

type Whitelist map[string][]string
