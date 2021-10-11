package http

import (
	"net/http"

	walletcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
)

type server struct {
	walletUsecase *walletcore.WalletUsecase
}

func New(walletUsecase *walletcore.WalletUsecase) *server {
	return &server{
		walletUsecase: walletUsecase,
	}
}

func (s *server) ListenAndServe() error {
	r := s.router()
	if err := http.ListenAndServe(":3000", r); err != nil {
		return err
	}
	return nil
}
