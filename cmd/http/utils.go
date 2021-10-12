package http

import (
	"net/http"
	"strings"
)

func sendCustomHttpError(w http.ResponseWriter, err error) {
	if strings.HasPrefix(err.Error(), "invalid") {
		writeBadRequest(w, err, http.StatusBadRequest)
	} else if strings.HasSuffix(err.Error(), "already_exists") {
		writeBadRequest(w, err, http.StatusConflict)
	} else if strings.HasSuffix(err.Error(), "not_found") {
		writeBadRequest(w, err, http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func writeBadRequest(w http.ResponseWriter, err error, status int) {
	if status == 0 {
		status = http.StatusBadRequest
	}

	w.WriteHeader(status)
	w.Write([]byte(err.Error()))
}
