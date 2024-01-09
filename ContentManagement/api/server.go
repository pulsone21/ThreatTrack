package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	*http.Server
	*mux.Router
	Db DB
}

type APIFunc func(http.ResponseWriter, *http.Request) (*ApiResponse, *ApiError)

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
	go func() {
		backendPort := os.Getenv("BACKEND_PORT")
		fmt.Println(fmt.Sprintf("Starting Backend Server on https://localhost:%s", backendPort))
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", backendPort), server))
	}()
	// Implement graceful shutdown.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down the server...")

	// Set a timeout for shutdown (for example, 5 seconds).
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
	log.Println("Server gracefully stopped")
}

func createHandleFunc(apiF APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := apiF(w, r)
		if err != nil {
			writeJSON(w, err.StatusCode, map[string]string{"RequestUrl": err.RequestUrl, "Message": err.Error()})
		}
		writeJSON(w, res.StatusCode, res)
	}
}

func writeJSON(w http.ResponseWriter, status int, val any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(val)
}