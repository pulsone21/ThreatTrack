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
	Db DB
}

type APIFunc func(http.ResponseWriter, *http.Request) error

func NewServer() *ApiServer {
	dbUsr := os.Getenv("MYSQLUSER")
	dbPw := os.Getenv("MYSQLPW")
	dbPort := os.Getenv("DB_PORT")
	dbIP := os.Getenv("DB_ADRESS")
	fmt.Printf("DB_USER: %s, DB_PW: %s, DB_PORT: %s, DB_IP: %s\n", dbUsr, dbPw, dbPort, dbIP)
	s := &ApiServer{
		Router: mux.NewRouter(),
	}

	s.Db = *setupDB(dbIP+":"+dbPort, dbUsr, dbPw, s)
	return s
}

func (server *ApiServer) Run() {
	backendPort := os.Getenv("BACKEND_PORT")
	fmt.Println(fmt.Sprintf("Starting Backend Server on https://localhost:%s", backendPort))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", backendPort), server))
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
