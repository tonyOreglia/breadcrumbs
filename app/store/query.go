package store

import (
	"fmt"

	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

type Note struct {
	Message         string  `json:"message" db:"note"`
	Latitude        float64 `json:"latitude" db:"st_x"`
	Longitude       float64 `json:"longitude" db:"st_y"`
	Altitude        float64 `json:"altitude" db:"st_z"`
	DateCreatedUnix uint64  `json:"date_created_unix" db:"date_created_unix"`
}

// SaveNote stores textual data at a single point location
func (s *Store) SaveNote(note string, latitude float64, longitude float64, altitude float64, dateCreatedUnix uint64) error {
	log.Infof("Saving note: '%s'", note)
	txn, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer txn.Rollback() // The rollback will be ignored if the tx has been committed later in the function.
	query := `INSERT INTO notes (note) VALUES ($1) ON CONFLICT DO NOTHING`
	_, err = txn.Exec(query, note)
	if err != nil {
		return err
	}

	query = fmt.Sprintf(`
		INSERT INTO breadcrumbs (data_type, data_id, geog, date_created_unix) VALUES ( 'note', (SELECT id from notes WHERE note=$1), ST_SetSRID(ST_MakePoint(%.12f, %.12f, %.1f), 4326), %d)
		`, latitude, longitude, altitude, dateCreatedUnix)
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

func (s *Store) RetrieveNotes(radiusInMeters int, lat float64, long float64) (error, []Note) {
	log.Println(fmt.Sprintf("Retrieving notes within %d meters of latitude(%.12f), longitude(%.12f)", radiusInMeters, lat, long))
	notes := []Note{}
	query := fmt.Sprintf(`
		SELECT n.note, ST_X(b.geog::geometry), ST_Y(geog::geometry), ST_Z(geog::geometry), b.date_created_unix FROM breadcrumbs as b LEFT JOIN notes as n ON b.data_id = n.id WHERE ST_DWithin(b.geog, ST_MakePoint(%.12f, %.12f), $1)
		`, lat, long)
	err := s.DB.Select(&notes, query, radiusInMeters)
	if err != nil {
		return err, notes
	}
	log.Println(fmt.Sprintf("Retrieve %d notes: %v: ", len(notes), notes))
	return nil, notes
}

func (s *Store) SaveNotes(notes []Note) error {
	log.Infof("Saving %d notes", len(notes))
	txn, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	saveNotesQuery := "INSERT INTO notes (note) VALUES "
	saveLocationQuery := "INSERT INTO breadcrumbs (data_type, data_id, geog, date_created_unix) VALUES "
	for index, note := range notes {
		saveNotesQuery = saveNotesQuery + fmt.Sprintf("('%s')", note.Message)

		saveLocationQuery = saveLocationQuery + fmt.Sprintf(
			"( 'note', (SELECT id from notes WHERE note='%s'), 'SRID=4326;POINTZ(%.12f %.12f %.1f)', %d)",
			note.Message,
			note.Latitude,
			note.Longitude,
			note.Altitude,
			note.DateCreatedUnix,
		)

		if index != (len(notes) - 1) {
			saveNotesQuery = saveNotesQuery + ", "
			saveLocationQuery = saveLocationQuery + ", "
		} else {
			saveNotesQuery = saveNotesQuery + " ON CONFLICT DO NOTHING;"
			saveLocationQuery = saveLocationQuery + ";"
		}
	}
	_, err = txn.Exec(saveNotesQuery)
	if err != nil {
		return err
	}
	_, err = txn.Exec(saveLocationQuery)
	if err != nil {
		return err
	}
	err = txn.Commit()
	if err != nil {
		return err
	}
	return nil
}
