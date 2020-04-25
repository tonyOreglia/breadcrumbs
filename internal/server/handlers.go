package server

import (
    "fmt"
    "errors"
    "io"
    "strings"
	"net/http"
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/golang/gddo/httputil/header"
)

// TO DO: refactor this by embedding Location
// TO DO: refactor this by adding proper json property namees (lowercase)
type Note struct {
    Note string
	Latitude float64
	Longitude float64
}

type Location struct {
    Latitude float64  `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Area struct {
    Location Location  `json:"location"`
    RadiusInMeters int `json:"radius_in_meters"`
}

// process a CSV payload of mobile numbers
func (s *Server) storeNoteHandler(w http.ResponseWriter, r *http.Request) {
	var n Note
	err := decodeJSONBody(w, r, &n)
    if err != nil {
        var mr *malformedRequest
        if errors.As(err, &mr) {
            http.Error(w, mr.msg, mr.status)
        } else {
            log.Println(err.Error())
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        }
        return
	}
	s.db.SaveNote(n.Note, n.Latitude, n.Longitude)
}

func (s *Server) getNotesHandler(w http.ResponseWriter, r *http.Request) {
	var a Area
	err := decodeJSONBody(w, r, &a)
    if err != nil {
        var mr *malformedRequest
        if errors.As(err, &mr) {
            http.Error(w, mr.msg, mr.status)
        } else {
            log.Println(err.Error())
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        }
        return
	}
    err, notes := s.db.RetrieveNotes(a.Location.Latitude, a.Location.Longitude, a.RadiusInMeters)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(notes)
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