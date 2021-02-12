package server

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Server handles http requests
type Server struct {
}

// NewServer creates a new instance of the server
func NewServer() *Server {
	server := Server{}
	return &server
}

// Run runs the server
func (s *Server) Run(port int) {
	router := mux.NewRouter()
	router.HandleFunc("/accept", handleAccept).Methods("POST")
	http.ListenAndServe(":"+strconv.Itoa(10000), router)
}
