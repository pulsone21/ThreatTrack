package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	*mux.Router
	ListenAddr string
	db         DB
}

type APIFunc func(http.ResponseWriter, *http.Request) error

func NewServer(listenAddr string) *ApiServer {
	return &ApiServer{
		Router:     mux.NewRouter(),
		ListenAddr: listenAddr,
	}
}

func (server *ApiServer) Run() {
	dbUsr := os.Getenv("MYSQLUSER")
	dbPw := os.Getenv("MYSQLPW")
	dbPort := os.Getenv("DBPORT")
	fmt.Printf("DB_USER: %s, DB_PW: %s, DB_PORT: %s\n", dbUsr, dbPw, dbPort)
	server.db = *setupDB("localhost:"+dbPort, dbUsr, dbPw, server)
	fmt.Println(fmt.Sprintf("Starting Backend Server on https://localhost:%s", server.ListenAddr))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.ListenAddr), server))
}

type ApiError struct {
	RequestUrl string
	Error      string `json:"error"`
}

type APIResponse struct {
	Data []interface{}
}

func createHandleFunc(apiF APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiF(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, ApiError{RequestUrl: r.RequestURI, Error: err.Error()})
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, val any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(val)
}
