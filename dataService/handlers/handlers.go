package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"dataservice/storage"
	"threattrack/entities"

	"github.com/gorilla/mux"
)

func CreateHandlers(r *mux.Router, s storage.Storage) {
	inc_sR := r.PathPrefix("/incidents").Subrouter()
	inc_sR.HandleFunc("", newHandleFunc(s.HandleGetAll)).Methods(http.MethodGet)
	inc_sR.HandleFunc("", newHandleFunc(s.HandleCreate)).Methods(http.MethodPost)
	inc_sR.HandleFunc("/{id}", newHandleFunc(s.HandleGetID)).Methods(http.MethodGet)
	inc_sR.HandleFunc("/{id}", newHandleFunc(s.HandleUpdate)).Methods(http.MethodPut)
	inc_sR.HandleFunc("/{id}", newHandleFunc(s.HandleDelete)).Methods(http.MethodDelete)
	inc_sR.HandleFunc("/query", newHandleFunc(s.HandleGetQuery)).Methods(http.MethodGet)

	it_sR := r.PathPrefix("/incidenttypes").Subrouter()
	it_sR.HandleFunc("", newHandleFunc(s.HandleGetAll)).Methods(http.MethodGet)
	it_sR.HandleFunc("", newHandleFunc(s.HandleCreate)).Methods(http.MethodPost)
	it_sR.HandleFunc("/{id}", newHandleFunc(s.HandleGetID)).Methods(http.MethodGet)
	it_sR.HandleFunc("/{id}", newHandleFunc(s.HandleUpdate)).Methods(http.MethodPut)
	it_sR.HandleFunc("/{id}", newHandleFunc(s.HandleDelete)).Methods(http.MethodDelete)

	usr_sR := r.PathPrefix("/users").Subrouter()
	usr_sR.HandleFunc("", newHandleFunc(s.HandleGetAll)).Methods(http.MethodGet)
	usr_sR.HandleFunc("", newHandleFunc(s.HandleCreate)).Methods(http.MethodPost)
	usr_sR.HandleFunc("/{id}", newHandleFunc(s.HandleGetID)).Methods(http.MethodGet)
	usr_sR.HandleFunc("/{id}", newHandleFunc(s.HandleUpdate)).Methods(http.MethodPut)
	usr_sR.HandleFunc("/{id}", newHandleFunc(s.HandleDelete)).Methods(http.MethodDelete)

	task_sR := r.PathPrefix("/tasks").Subrouter()
	task_sR.HandleFunc("", newHandleFunc(s.HandleGetAll)).Methods(http.MethodGet)
	task_sR.HandleFunc("", newHandleFunc(s.HandleCreate)).Methods(http.MethodPost)
	task_sR.HandleFunc("/{id}", newHandleFunc(s.HandleGetID)).Methods(http.MethodGet)
	task_sR.HandleFunc("/{id}", newHandleFunc(s.HandleUpdate)).Methods(http.MethodPut)
	task_sR.HandleFunc("/{id}", newHandleFunc(s.HandleDelete)).Methods(http.MethodDelete)
	task_sR.HandleFunc("/query", newHandleFunc(s.HandleGetQuery)).Methods(http.MethodGet)
}

func newHandleFunc(apiF entities.APIFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "uri", r.RequestURI)
		fmt.Println(r.URL.Path)
		ctx = context.WithValue(ctx, "entity", strings.Split(r.URL.Path, "/")[1])
		fmt.Println("----------New API Request-------------")
		res, err := apiF(ctx, w, r)
		if err != nil {
			Respond(w, err.StatusCode, map[string]string{"RequestUrl": err.RequestUrl, "Message": err.Error()})
			fmt.Println("----------API Request finished with error-------------")
			return
		}
		if res != nil {
			Respond(w, res.StatusCode, res)
			fmt.Println("----------API Request finished with no result-------------")
			return
		}
		Respond(w, http.StatusInternalServerError, map[string]string{"RequestUrl": r.RequestURI, "Message": "error and response are nil"})
		fmt.Println("----------API Request finished successfully-------------")
	}
}

func Respond(w http.ResponseWriter, status int, val any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(val)
}
