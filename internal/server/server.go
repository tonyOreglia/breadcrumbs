package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tonyOreglia/breadcrumbs/store"
	log "github.com/sirupsen/logrus"
)

var (
	url  = "http://localhost"
	port = 80
)

// Server defines a HTTP Server
type Server struct {
	r  *mux.Router
	db *store.Store
}

// New returns HTTP Server configured for localhost port 80
func New() *Server {
	log.Info("Starting server...")
	server := new(Server)
	server.db = store.New("postgresql://toreglia:anthony@localhost/breadcrumbs?sslmode=disable", 2)
	server.r = mux.NewRouter()
	server.r.HandleFunc("/note", server.storeNoteHandler).Methods("POST")
	server.r.HandleFunc("/getNotes", server.getNotesHandler).Methods("POST")
	return server
}

// Start starts the server
func (s *Server) Start() error {
	return http.ListenAndServe(":80", s.r)
}
