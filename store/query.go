package store

import (
	"fmt"
	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

type Breadcrumb struct {
	DataType int `json:"data_type"`
	DataId   int `json:"data_id"`
	Location int `json:"geog"`
}

// SaveNote stores textual data at a single point location
func (s *Store) SaveNote(note string, latitude float64, longitude float64) error {
	log.Infof("Saving note: '%s'", note)
	txn, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer txn.Rollback() // The rollback will be ignored if the tx has been committed later in the function.
	query := `INSERT INTO notes (note) VALUES ($1)`
	_, err = txn.Exec(query, note)
	if err != nil {
		return err
	}

	query = fmt.Sprintf(`INSERT INTO breadcrumbs (data_type, data_id, geog) VALUES ( 'notes', (SELECT id from notes WHERE note=$1), 'SRID=4326;POINT(%.6f %.6f)')`, latitude, longitude)
	_, err = txn.Exec(query, note)
	if err != nil {
		return err
	}
	err = txn.Commit()
	if err != nil {
		return err
	}
	return nil
}
