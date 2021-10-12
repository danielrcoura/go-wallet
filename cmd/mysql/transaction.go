package mysql

import (
	"database/sql"
	"fmt"

	wcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
)

type transactionMysql struct {
	db *sql.DB
}

func NewTransactionMysql(db *sql.DB) *transactionMysql {
	return &transactionMysql{
		db: db,
	}
}

func (t *transactionMysql) FetchByWallet(walletID int) ([]wcore.Transaction, error) {
	r, err := t.db.Query(`SELECT id, ticker, operation, quantity, price, date 
	                      FROM transactions
						  WHERE wallet_id=?`, walletID)
	if err != nil {
		return nil, err
	}

	return RowsToTransactions(r)
}

func (t *transactionMysql) Store(walletID int, transaction wcore.Transaction) error {
	data := []interface{}{
		walletID,
		transaction.Ticker,
		transaction.Operation,
		transaction.Quantity,
		transaction.Price,
		transaction.Date,
	}

	_, err := t.db.Exec(`INSERT INTO transactions (wallet_id, ticker, operation, quantity, price, date) 
						 VALUES (?, ?, ?, ?, ?, ?)`, data...)
	if err != nil {
		return err
	}

	return nil
}

func (t *transactionMysql) Update() {
	fmt.Println("transaction_repository: update")
}

func (t *transactionMysql) Delete() {
	fmt.Println("transaction_repository: delete")
}
