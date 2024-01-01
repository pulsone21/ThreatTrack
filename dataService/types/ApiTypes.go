package types

import (
	"context"
	"net/http"
)

type ApiError struct {
	error
	StatusCode int
	RequestUrl string
}
type ScanFunc func(dest ...any) error

type ApiResponse struct {
	StatusCode int
	RequestUrl string
	Data       interface{}
}

type APIFunction func(context.Context, http.ResponseWriter, *http.Request) (*ApiResponse, *ApiError)

func InternalServerError(err error, uri string) *ApiError {
	return NewApiError(http.StatusInternalServerError, uri, err)
}

func BadRequestError(err error, uri string) *ApiError {
	return NewApiError(http.StatusBadRequest, uri, err)
}
func NotFoundError(err error, uri string) *ApiError {
	return NewApiError(http.StatusNotFound, uri, err)
}

func NewApiError(status int, uri string, err error) *ApiError {
	return &ApiError{
		error:      err,
		StatusCode: status,
		RequestUrl: uri,
	}
}

func NewApiResponse(statusCode int, uri string, data interface{}) *ApiResponse {
	return &ApiResponse{
		StatusCode: statusCode,
		RequestUrl: uri,
		Data:       data,
	}
}
