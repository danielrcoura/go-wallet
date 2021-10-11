package http

import (
	"encoding/json"
	"net/http"

	"github.com/danielrcoura/go-wallet/cmd/walleterror"
)

type userReq struct {
	Name string `json:"name"`
}

func (s *server) fetchWallets(w http.ResponseWriter, r *http.Request) {
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

func (s *server) storeWallet(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u userReq
	if err := decoder.Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := s.walletUsecase.Store(u.Name); err != nil {
		switch err.Error() {
		case walleterror.ErrInvalidWalletName.Error():
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(walleterror.ErrInvalidWalletName.Error()))
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
}
