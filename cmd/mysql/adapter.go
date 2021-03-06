package mysql

import (
	"database/sql"

	wcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
)

func RowsToWallets(r *sql.Rows) ([]wcore.SimpleWallet, error) {
	wallets := []wcore.SimpleWallet{}

	for r.Next() {
		w := wcore.SimpleWallet{}
		err := r.Scan(&w.ID, &w.Name)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, w)
	}

	return wallets, nil
}

func RowsToTransactions(r *sql.Rows) ([]wcore.Transaction, error) {
	transactions := []wcore.Transaction{}

	for r.Next() {
		t := wcore.Transaction{}
		err := r.Scan(&t.ID, &t.Ticker, &t.Quantity, &t.Price, &t.Date)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}
