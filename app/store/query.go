package store

import (
	"fmt"
	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

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

	query = fmt.Sprintf(`
		INSERT INTO breadcrumbs (data_type, data_id, geog) VALUES ( 'notes', (SELECT id from notes WHERE note=$1), 'SRID=4326;POINT(%.6f %.6f)')
		`, latitude, longitude)
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

type Note struct {
	Note string         `json:"note" db:"note"`
	Latitude float64  `json:"latitude" db:"st_x"`
	Longitude float64 `json:"longitude" db:"st_y"`
}

func (s *Store) RetrieveNotes(radiusInMeters int, lat float64, long float64) (error, []Note) {
	notes := []Note{}
	query := fmt.Sprintf(`
		SELECT n.note, ST_X(b.geog::geometry), ST_Y(geog::geometry) FROM breadcrumbs as b LEFT JOIN notes as n ON b.data_id = n.id WHERE ST_DWithin(b.geog, ST_MakePoint(%.6f, %.6f), $1)
		`, lat, long)
	s.DB.Select(&notes, query, radiusInMeters)
	return nil, notes
}