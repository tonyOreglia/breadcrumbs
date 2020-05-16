package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/lib/pq"
	"github.com/golang/gddo/httputil/header"
	log "github.com/sirupsen/logrus"

	"github.com/tonyOreglia/breadcrumbs/store"
)

func (s *Server) storeNoteHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody store.Note
	err := decodeJSONBody(w, r, &requestBody)
	if err != nil {
		handleHttpError(err, w, r)
		return
	}
	err = s.db.SaveNote(requestBody.Message, requestBody.Latitude, requestBody.Longitude)
	if err != nil {
		handleHttpError(err, w, r)
	}
}

func (s *Server) storeNotesHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody []store.Note
	err := decodeJSONBody(w, r, &requestBody)
	if err != nil {
		handleHttpError(err, w, r)
		return
	}
	s.db.SaveNotes(requestBody)
}

func (s *Server) getNotesHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		RadiusInMeters int     `json:"radius_in_meters"`
		Latitude       float64 `json:"latitude"`
		Longitude      float64 `json:"longitude"`
	}
	err := decodeJSONBody(w, r, &requestBody)
	if err != nil {
		handleHttpError(err, w, r)
		return
	}
	err, notes := s.db.RetrieveNotes(requestBody.RadiusInMeters, requestBody.Latitude, requestBody.Longitude)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(notes)
}

func handleHttpError(err error, w http.ResponseWriter, r *http.Request) {
	var mr *malformedRequest
	if errors.As(err, &mr) {
		http.Error(w, mr.msg, mr.status)
		return;
	}
	if err, ok := err.(*pq.Error); ok {
		http.Error(w, err.Message, 400)
	}
	log.Println(err.Error())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

type malformedRequest struct {
	status int
	msg    string
}

func (mr *malformedRequest) Error() string {
	return mr.msg
}

func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			return &malformedRequest{status: http.StatusUnsupportedMediaType, msg: msg}
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return &malformedRequest{status: http.StatusRequestEntityTooLarge, msg: msg}

		default:
			return err
		}
	}

	if dec.More() {
		msg := "Request body must only contain a single JSON object"
		return &malformedRequest{status: http.StatusBadRequest, msg: msg}
	}

	return nil
}
