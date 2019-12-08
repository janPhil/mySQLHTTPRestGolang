package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/janPhil/mySQLHTTPRestGolang/database"
	"github.com/janPhil/mySQLHTTPRestGolang/types"
)

// Server is a struct which contains all dependencies for this microservice
type Server struct {
	Router *mux.Router
	db     *sql.DB
}

// NewServer returns an instance of the server with all dependencies.
// The served routes can be found in the routes.go file
func NewServer() (*Server, error) {
	db, err := database.NewDB()
	if err != nil {
		return nil, err
	}
	s := &Server{
		Router: mux.NewRouter(),
		db:     db,
	}
	s.createRoutes()
	fmt.Println("Server created")
	return s, nil
}

func (s *Server) createRoutes() {
	s.Router.HandleFunc("/", s.getIndex())
	s.Router.HandleFunc("/all", s.getAll())
}

func (s *Server) getIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Welcome")
	}
}

func (s *Server) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := types.AllEmployees(s.db)
		if err != nil {
			json.NewEncoder(w).Encode("Failed to get Employees")
		}
		json.NewEncoder(w).Encode(res)
	}
}

// StartServer starts the new created Server
func (s *Server) StartServer() {
	fmt.Println("Server started")
	http.ListenAndServe(":8080", s.Router)
	defer s.StopServer()
}

// StopServer stops the server and closes the connection to the database
func (s *Server) StopServer() {
	fmt.Println("Shutting down server...")
	fmt.Println("closing database...")
	s.db.Close()
}
