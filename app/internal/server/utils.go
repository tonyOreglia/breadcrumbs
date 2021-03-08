package server

import (
	"database/sql"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/tonyOreglia/breadcrumbs/app/store"
)

func handleError(w http.ResponseWriter, err error, code int) {
	log.Error(err)
	errJSON := jsonError{Msg: err.Error()}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errJSON)
}

func transformNotesForDb(notes []Note) []store.Note {
	var dbNotes []store.Note
	for _, note := range notes {
		dbNotes = append(dbNotes, transformForDb(note))
	}
	return dbNotes
}

func transformForDb(note Note) store.Note {
	return store.Note{
		Id: note.Id,
		Message: note.Message,
		Latitude: note.Latitude,
		Longitude: note.Longitude,
		Altitude: note.Altitude,
		UserName: sql.NullString{String: note.UserName, Valid: true},
		Timestamp: note.Timestamp,
	}
}

func transformNotesForClient(notes []store.Note) []Note {
	var dbNotes []Note
	for _, note := range notes {
		dbNotes = append(dbNotes, transformForClient(note))
	}
	return dbNotes
}


func transformForClient(note store.Note) Note {
	return Note{
		Id: note.Id,
		Message: note.Message,
		Latitude: note.Latitude,
		Longitude: note.Longitude,
		Altitude: note.Altitude,
		UserName: note.UserName.String,
		Timestamp: note.Timestamp,
	}
}