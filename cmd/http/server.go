package http

import (
	"net/http"

	wcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
)

type server struct {
	swalletUsecase     *wcore.SimpleWalletUsecase
	walletUsecase      *wcore.WalletUsecase
	transactionUsecase *wcore.TransactionUsecase
	coinUsecase        *wcore.CoinUsecase
}

func New(
	swalletUsecase *wcore.SimpleWalletUsecase,
	transactionUsecase *wcore.TransactionUsecase,
	coinUsecase *wcore.CoinUsecase,
	walletUsecase *wcore.WalletUsecase,
) *server {
	return &server{
		swalletUsecase:     swalletUsecase,
		transactionUsecase: transactionUsecase,
		coinUsecase:        coinUsecase,
		walletUsecase:      walletUsecase,
	}
}

func (s *server) ListenAndServe() error {
	r := s.router()
	if err := http.ListenAndServe(":3000", r); err != nil {
		return err
	}
	return nil
}
