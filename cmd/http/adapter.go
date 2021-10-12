package http

import (
	"strings"
	"time"

	wcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
)

func walletReqToWallet(wReq walletReq) wcore.Wallet {
	return wcore.Wallet{
		Name: wReq.Name,
	}
}

func transactionReqToTransaction(tReq transactionReq) (wcore.Transaction, error) {
	t := wcore.Transaction{
		Ticker:    tReq.Ticker,
		Operation: wcore.StringToOperation(tReq.Operation),
		Quantity:  tReq.Quantity,
		Price:     tReq.Price,
		Date:      nil,
	}

	if strings.TrimSpace(tReq.Date) != "" {
		layout := "2006-01-02T15:04:05Z"
		date, err := time.Parse(layout, tReq.Date)
		if err != nil {
			return t, wcore.ErrInvalidTransactionDate
		}
		t.Date = &date
	}

	return t, nil
}
