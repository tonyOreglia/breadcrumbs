package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Store defines a postgres data store
type Store struct {
	DB *sqlx.DB
}

// Close connection with the db
func (s *Store) Close() {
	if s.DB != nil {
		err := s.DB.Close()
		if err != nil {
			log.Error(errors.Wrap(err, "failed to close connection"))
		}
	}
}

// New returns the postgres implementation of a data store
func New(connString string, maxDBConns int) *Store {
	store := &Store{}
	store.DB = sqlx.MustConnect("postgres", connString)
	store.DB.SetMaxOpenConns(maxDBConns)
	return store
}
