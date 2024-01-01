package handlers

import (
	"context"
	"data-service/storage"
	"data-service/types"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func CreateHandlers(r *mux.Router, s storage.Storage) {
	inc_sR := r.PathPrefix("/incidents").Subrouter()
	inc_sR.HandleFunc("", newHandleFunc(s.HandleGetAll)).Methods(http.MethodGet)
	inc_sR.HandleFunc("", newHandleFunc(s.HandleCreate)).Methods(http.MethodPost)
	inc_sR.HandleFunc("/query", newHandleFunc(s.HandleGetQuery)).Methods(http.MethodGet)
	inc_sR.HandleFunc("/{id}", newHandleFunc(s.HandleGetID)).Methods(http.MethodGet)
	inc_sR.HandleFunc("/{id}", newHandleFunc(s.HandleUpdate)).Methods(http.MethodPut)
	inc_sR.HandleFunc("/{id}", newHandleFunc(s.HandleDelete)).Methods(http.MethodDelete)

}

func newHandleFunc(apiF types.APIFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "uri", r.RequestURI)
		fmt.Println(r.URL.Path)
		ctx = context.WithValue(ctx, "entity", strings.Split(r.URL.Path, "/")[1])

		res, err := apiF(ctx, w, r)
		if err != nil {
			Respond(w, err.StatusCode, map[string]string{"RequestUrl": err.RequestUrl, "Message": err.Error()})
			return
		}
		if res != nil {
			Respond(w, res.StatusCode, res)
			return
		}
		Respond(w, http.StatusInternalServerError, map[string]string{"RequestUrl": r.RequestURI, "Message": "error and response are nil"})
	}
}

func Respond(w http.ResponseWriter, status int, val any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(val)
}