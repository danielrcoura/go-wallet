package mysql

import (
	"database/sql"

	walletcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
)

func RowsToWallets(r *sql.Rows) ([]walletcore.Wallet, error) {
	wallets := []walletcore.Wallet{}

	for r.Next() {
		w := walletcore.Wallet{}
		err := r.Scan(&w.ID, &w.Name)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, w)
	}

	return wallets, nil
}

func RowsToTransactions(r *sql.Rows) ([]walletcore.Transaction, error) {
	transactions := []walletcore.Transaction{}

	// TODO: change Operation to enum
	// TODO: change Date to time
	for r.Next() {
		t := walletcore.Transaction{}
		err := r.Scan(&t.ID, nil, &t.Ticker, &t.Operation, &t.Quantity, &t.Price, &t.Date)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}
