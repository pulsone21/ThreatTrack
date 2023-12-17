package gui

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
	Port string
}

type PageHandler interface {
	createHandles(*Server)
}

func CreateServer(port, backendAdress string) *Server {
	server := &Server{
		Router: mux.NewRouter(),
		Port:   port,
	}

	CreateIncidentHandler(server, backendAdress)
	CreateIndicatorHandler(server, backendAdress)
	server.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static/"))))
	fmt.Println("FrontEnd Created")
	return server
}

func (s *Server) Run() {
	fmt.Printf("Serving Webserver at https://localhost:%s", s.Port)
	panic(http.ListenAndServe(fmt.Sprintf(":%s", s.Port), s))
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
