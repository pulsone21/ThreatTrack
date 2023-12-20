package api

import (
	"net/http"
)

type ApiError struct {
	// TODO Redesign the API Response with reasonable expectable return Objects
	error
	StatusCode int
	RequestUrl string
}

type ApiResponse struct {
	StatusCode int
	RequestUrl string
	Data       interface{}
}

func InternalServerError(err error, uri string) *ApiError {
	return &ApiError{
		error:      err,
		StatusCode: http.StatusInternalServerError,
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
