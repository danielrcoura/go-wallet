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
		sendCustomHttpError(w, err)
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
		sendCustomHttpError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *server) updateWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars[WALLET_ID])
	if err != nil {
		sendCustomHttpError(w, wcore.ErrWalletNotFound)
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
		sendCustomHttpError(w, err)
		return
	}
}

func (s *server) deleteWallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars[WALLET_ID])
	if err != nil {
		sendCustomHttpError(w, wcore.ErrWalletNotFound)
		return
	}

	if err := s.walletUsecase.Delete(id); err != nil {
		sendCustomHttpError(w, err)
	}
}
