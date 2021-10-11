package http

import "net/http"

func (s *server) router() {
	http.HandleFunc("/wallets", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			s.fetchWallets(w, r)
		case http.MethodPost:
			s.storeWallet(w, r)
		}
	})
}
