package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const WALLET_ID = "id"

func (s *server) router() *mux.Router {
	r := mux.NewRouter()

	sub := r.PathPrefix("/wallets").Subrouter()
	sub.HandleFunc("", s.fetchWallets).Methods(http.MethodGet)
	sub.HandleFunc("", s.storeWallet).Methods(http.MethodPost)
	sub.HandleFunc(fmt.Sprintf("/{%s}", WALLET_ID), s.updateWallet).Methods(http.MethodPatch)
	sub.HandleFunc(fmt.Sprintf("/{%s}", WALLET_ID), s.deleteWallet).Methods(http.MethodDelete)

	// sub = r.PathPrefix("/wallets/{walletID}/transactions").Subrouter()
	// sub.HandleFunc("/", s.fetchWallets).Methods(http.MethodGet)
	// sub.HandleFunc("/", s.storeWallet).Methods(http.MethodPost)
	// sub.HandleFunc("/{id}", s.fetchWallets).Methods(http.MethodPatch)
	// sub.HandleFunc("/{id}", s.fetchWallets).Methods(http.MethodDelete)

	return r
}
