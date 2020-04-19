package server

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func handleError(w http.ResponseWriter, err error, code int) {
	log.Error(err)
	errJSON := jsonError{Msg: err.Error()}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errJSON)
}
