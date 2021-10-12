package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	wcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
	"github.com/gorilla/mux"
)

type walletReq struct {
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
	var wReq walletReq
	if err := decoder.Decode(&wReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := s.walletUsecase.Store(wReq.Name); err != nil {
		switch err.Error() {
		case wcore.ErrInvalidWalletName.Error():
			WriteBadRequestMsg(w, wcore.ErrInvalidWalletName)
		case wcore.ErrWalletAlreadyExists.Error():
			WriteBadRequestMsg(w, wcore.ErrWalletAlreadyExists)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *server) updateWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars[WALLET_ID])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var wReq walletReq
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&wReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	wallet := walletReqToWallet(wReq)

	if err := s.walletUsecase.Update(id, wallet); err != nil {
		switch err.Error() {
		case wcore.ErrInvalidWalletName.Error():
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(wcore.ErrInvalidWalletName.Error()))
		case wcore.ErrWalletAlreadyExists.Error():
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(wcore.ErrWalletAlreadyExists.Error()))
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
}

func (s *server) deleteWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars[WALLET_ID])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := s.walletUsecase.Delete(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
