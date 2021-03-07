package store

import (
	"fmt"
	"time"

	"database/sql"

	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

type Note struct {
	Id 				int 	`json:"id" db:"id"`
	Message         string  `json:"message" db:"note"`
	Latitude        float64 `json:"latitude" db:"st_y"`
	Longitude       float64 `json:"longitude" db:"st_x"`
	Altitude        float64 `json:"altitude" db:"st_z"`
	UserName		sql.NullString 	`json:"userName" db:"full_name"`
	Timestamp 		time.Time  `json:"timestamp" db:"ts"`
}

// SaveNote stores textual data at a single point location
func (s *Store) SaveNote(note string, latitude float64, longitude float64, altitude float64, userName sql.NullString) error {
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

	var id int
	query = `INSERT INTO bc_user (full_name) VALUES ($1) ON CONFLICT(full_name) DO UPDATE SET full_name=EXCLUDED.full_name RETURNING id`
	err = txn.QueryRow(query, userName).Scan(&id)
	if err != nil {
		return err
	}

	query = fmt.Sprintf(`
		INSERT INTO breadcrumbs (data_type, data_id, geog, user_id) VALUES ( 'note', (SELECT id from notes WHERE note=$1), ST_SetSRID(ST_MakePoint(%.12f, %.12f, %.1f), 4326), $2)
		`, longitude, latitude, altitude)
	fmt.Printf(query)
	_, err = txn.Exec(query, note, id)
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
		SELECT b.id, n.note, ST_Y(geog::geometry), ST_X(b.geog::geometry), ST_Z(geog::geometry), b.ts, bc_user.full_name
		FROM breadcrumbs as b 
		LEFT JOIN notes as n ON b.data_id = n.id
		LEFT JOIN bc_user ON bc_user.id = b.user_id
		WHERE ST_DWithin(b.geog, ST_MakePoint(%.12f, %.12f), $1)
		`, long, lat)
	err := s.DB.Select(&notes, query, radiusInMeters)
	if err != nil {
		return err, notes
	}
	log.Println(fmt.Sprintf("Retrieve %d notes: %v: ", len(notes), notes))
	return nil, notes
}

func (s *Store) RetrieveAllNotes() (error, []Note) {
	log.Println("Retrieving all notes")
	notes := []Note{}
	query := fmt.Sprintf(`
		SELECT b.id, n.note, ST_Y(geog::geometry), ST_X(b.geog::geometry), ST_Z(geog::geometry), b.ts, bc_user.full_name
		FROM breadcrumbs as b 
		LEFT JOIN notes as n ON b.data_id = n.id
		LEFT JOIN bc_user ON bc_user.id = b.user_id
		WHERE n.note IS NOT NULL
		`)
	err := s.DB.Select(&notes, query)
	if err != nil {
		return err, notes
	}
	log.Println(fmt.Sprintf("Retrieved %d notes", len(notes)))
	return nil, notes
}

// does not currently support saving user ID for each note
func (s *Store) SaveNotes(notes []Note) error {
	log.Infof("Saving %d notes", len(notes))
	txn, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer txn.Rollback()

	saveNotesQuery := "INSERT INTO notes (note) VALUES "
	saveLocationQuery := "INSERT INTO breadcrumbs (data_type, data_id, geog) VALUES "
	for index, note := range notes {
		saveNotesQuery = saveNotesQuery + fmt.Sprintf("('%s')", note.Message)

		saveLocationQuery = saveLocationQuery + fmt.Sprintf(
			"( 'note', (SELECT id from notes WHERE note='%s'), 'SRID=4326;POINTZ(%.12f %.12f %.1f)')",
			note.Message,
			note.Longitude,
			note.Latitude,
			note.Altitude,
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
