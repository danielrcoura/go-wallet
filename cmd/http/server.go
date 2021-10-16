package http

import (
	"net/http"

	wcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
)

type server struct {
	walletUsecase      *wcore.WalletUsecase
	transactionUsecase *wcore.TransactionUsecase
	coinUsecase        *wcore.CoinUsecase
}

func New(
	walletUsecase *wcore.WalletUsecase,
	transactionUsecase *wcore.TransactionUsecase,
	coinUsecase *wcore.CoinUsecase,
) *server {
	return &server{
		walletUsecase:      walletUsecase,
		transactionUsecase: transactionUsecase,
		coinUsecase:        coinUsecase,
	}
}

func (s *server) ListenAndServe() error {
	r := s.router()
	if err := http.ListenAndServe(":3000", r); err != nil {
		return err
	}
	return nil
}
