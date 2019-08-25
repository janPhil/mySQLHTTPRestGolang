package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/janPhil/mySQLHTTPRestGolang/models"
)

// Server is a struct which contains all dependencies for this microservice
type Server struct {
	Router *mux.Router
	db     *sql.DB
}

// NewServer returns an instance of the server with all dependencies.
// The served routes can be found in the routes.go file
func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		db:     models.NewDB(),
	}
	s.routes()
	fmt.Println("Server created")
	return s
}

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Welcome")
	}
}

func (s *Server) handleAll(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := models.AllEmployees(db)
		if err != nil {
			log.Fatalf("Could not receive employees %v", err)
		}
		json.NewEncoder(w).Encode(res)
	}
}

// StartServer starts the new created Server
func (s *Server) StartServer() {
	fmt.Println("Server started")
	http.ListenAndServe(":8080", s.Router)
}

// StopServer stops the server and closes the connection to the database
func (s *Server) StopServer() {
	s.db.Close()
}
