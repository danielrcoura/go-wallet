package http

import "net/http"

func WriteBadRequest(w http.ResponseWriter, err error, status int) {
	if status == 0 {
		status = http.StatusBadRequest
	}

	w.WriteHeader(status)
	w.Write([]byte(err.Error()))
}
