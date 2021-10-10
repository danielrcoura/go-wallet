package http

import (
	"net/http"

	"encoding/json"
)

func (s *server) fetchWallets(w http.ResponseWriter, req *http.Request) {
	wallets, err := s.walletUsecase.Fetch()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(wallets)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(json)
}
