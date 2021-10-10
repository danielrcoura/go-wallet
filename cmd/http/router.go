package http

import "net/http"

func (s *server) router() {
	http.HandleFunc("/wallets", func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			s.fetchWallets(res, req)
		}
	})
}
