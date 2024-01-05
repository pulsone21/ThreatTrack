package api

import (
	"dataservice/handlers"
	"dataservice/storage"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddress string, storage storage.Storage) *Server {
	return &Server{
		Router:     mux.NewRouter(),
		listenAddr: listenAddress,
		store:      storage,
	}
}

func (s *Server) Run() error {
	handlers.CreateHandlers(s.Router, s.store)
	return http.ListenAndServe(fmt.Sprintf(":%s", s.listenAddr), s)
}
