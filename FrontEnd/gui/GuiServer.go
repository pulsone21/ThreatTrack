package gui

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
	adress string
}

type PageHandler interface {
	createHandles(*Server)
}

func CreateServer(address, backendAdress string) *Server {
	fmt.Println("Creating new Webserver")
	server := &Server{
		Router: mux.NewRouter(),
		adress: address,
	}

	CreateIncidentHandler(server, backendAdress)
	CreateIndicatorHandler(server, backendAdress)
	server.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static/"))))
	fmt.Println("Webserver Created")
	return server
}

func (s *Server) RunServer() {
	fmt.Printf("Serving Webserver at http://%s", s.adress)
	if err := http.ListenAndServe(s.adress, s); err != nil {
		log.Fatal(err)
	}
}

type APIFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func writeHTML(w http.ResponseWriter, status int, val any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(val)
}

func (s *Server) createHandleFunc(apiF APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiF(w, r); err != nil {
			writeHTML(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
