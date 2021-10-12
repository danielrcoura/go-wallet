package http

import (
	"time"

	wcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
)

func walletReqToWallet(wReq walletReq) wcore.Wallet {
	return wcore.Wallet{
		Name: wReq.Name,
	}
}

func transactionReqToTransaction(tReq transactionReq) (wcore.Transaction, error) {
	layout := "2006-01-02T15:04:05.000Z"
	date, err := time.Parse(layout, tReq.Date)
	if err != nil {
		return wcore.Transaction{}, err
	}

	return wcore.Transaction{
		Ticker:    tReq.Ticker,
		Operation: wcore.StringToOperation(tReq.Operation),
		Quantity:  tReq.Quantity,
		Price:     tReq.Price,
		Date:      date,
	}, nil
}
