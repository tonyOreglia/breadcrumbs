package server

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/tonyOreglia/breadcrumbs/config"
	"github.com/tonyOreglia/breadcrumbs/store"
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
func New(config *config.Config) *Server {
	log.Info("Starting server")
	log.Info(config)
	server := new(Server)
	server.db = store.New(store.NewStoreParams{
		Host:       config.DBHost,
		Port:       config.DBPort,
		User:       config.DBUser,
		Password:   config.DBPassword,
		DBname:     config.DBName,
		MaxDBConns: config.MaxDBConns,
	})
	server.r = mux.NewRouter()
	server.r.HandleFunc("/note", server.storeNoteHandler).Methods("POST")
	server.r.HandleFunc("/notes", server.storeNotesHandler).Methods("POST")
	server.r.HandleFunc("/notes", server.getNotesHandler).Methods("GET")
	return server
}

// Start starts the server
func (s *Server) Start() error {
	return http.ListenAndServe(":80", s.r)
}
