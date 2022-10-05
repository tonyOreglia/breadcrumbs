package store

import (
	"fmt"

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

type NewStoreParams struct {
	Host       string
	Port       int
	User       string
	Password   string
	DBname     string
	MaxDBConns int
}

// New returns the postgres implementation of a data store
func New(params NewStoreParams) *Store {
	store := &Store{}
	log.Infof("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	params.Host, params.Port, params.User, params.Password, params.DBname)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		params.Host, params.Port, params.User, params.Password, params.DBname)
	var err error
	store.DB, err = sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	store.DB.SetMaxOpenConns(params.MaxDBConns)
	return store
}
