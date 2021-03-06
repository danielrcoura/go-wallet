package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const WALLET_ID = "w_id"
const TRANSACTION_ID = "t_id"

func (s *server) router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/coins", s.getCoins)

	sub := r.PathPrefix("/wallets").Subrouter()
	sub.HandleFunc("", s.fetchWallets).Methods(http.MethodGet)
	sub.HandleFunc("", s.storeWallet).Methods(http.MethodPost)
	sub.HandleFunc(fmt.Sprintf("/{%s}", WALLET_ID), s.updateWallet).Methods(http.MethodPatch)
	sub.HandleFunc(fmt.Sprintf("/{%s}", WALLET_ID), s.deleteWallet).Methods(http.MethodDelete)

	sub = r.PathPrefix(fmt.Sprintf("/wallets/{%s}/transactions", WALLET_ID)).Subrouter()
	sub.HandleFunc("", s.fetchTransactions).Methods(http.MethodGet)
	sub.HandleFunc("", s.storeTransaction).Methods(http.MethodPost)
	sub.HandleFunc(fmt.Sprintf("/{%s}", TRANSACTION_ID), s.updateTransaction).Methods(http.MethodPatch)
	sub.HandleFunc(fmt.Sprintf("/{%s}", TRANSACTION_ID), s.deleteTransaction).Methods(http.MethodDelete)

	return r
}
