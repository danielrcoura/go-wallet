package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	wcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
	"github.com/gorilla/mux"
)

type transactionReq struct {
	Ticker    string  `json:"ticker"`
	Operation string  `json:"operation"`
	Quantity  float64 `json:"quantity"`
	Price     float64 `json:"price"`
	Date      string  `json:"date"`
}

func (s *server) fetchTransactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	wID, err := strconv.Atoi(vars[WALLET_ID])
	if err != nil {
		sendCustomHttpError(w, wcore.ErrWalletNotFound)
		return
	}

	transactions, err := s.transactionUsecase.FetchByWallet(wID)
	if err != nil {
		sendCustomHttpError(w, err)
		return
	}

	json, err := json.Marshal(transactions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func (s *server) storeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	wID, err := strconv.Atoi(vars[WALLET_ID])
	if err != nil {
		sendCustomHttpError(w, wcore.ErrWalletNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var tReq transactionReq
	if err := decoder.Decode(&tReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t, err := transactionReqToTransaction(tReq)
	if err != nil {
		sendCustomHttpError(w, err)
		return
	}

	if err := s.transactionUsecase.Store(wID, t); err != nil {
		sendCustomHttpError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *server) updateTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tID, err := strconv.Atoi(vars[TRANSACTION_ID])
	if err != nil {
		sendCustomHttpError(w, wcore.ErrTransactionNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var tReq transactionReq
	if err := decoder.Decode(&tReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t, err := transactionReqToTransaction(tReq)
	if err != nil {
		sendCustomHttpError(w, err)
		return
	}

	if err := s.transactionUsecase.Update(tID, t); err != nil {
		sendCustomHttpError(w, err)
		return
	}
}

func (s *server) deleteTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars[TRANSACTION_ID])
	if err != nil {
		sendCustomHttpError(w, wcore.ErrTransactionNotFound)
		return
	}

	if err := s.transactionUsecase.Delete(id); err != nil {
		sendCustomHttpError(w, err)
		return
	}
}
