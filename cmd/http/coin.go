package http

import (
	"encoding/json"
	"net/http"
)

func (s *server) getCoins(w http.ResponseWriter, r *http.Request) {
	coins, err := s.coinUsecase.GetCoins()
	if err != nil {
		sendCustomHttpError(w, err)
		return
	}

	json, err := json.Marshal(coins)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(json)
}
