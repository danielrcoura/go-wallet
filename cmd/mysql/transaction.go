package infra

import (
	"database/sql"
	"fmt"

	adapter "github.com/danielrcoura/go-wallet/cmd/adapters"
	walletcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
)

type transactionMysql struct {
	db *sql.DB
}

func NewTransactionMysql(db *sql.DB) *transactionMysql {
	return &transactionMysql{
		db: db,
	}
}

func (t *transactionMysql) FetchByWallet(w int) ([]walletcore.Transaction, error) {
	stmt, err := t.db.Prepare("SELECT * FROM transactions WHERE wallet_id=?")
	if err != nil {
		return nil, err
	}

	r, err := stmt.Query(w)
	if err != nil {
		return nil, err
	}

	return adapter.RowsToTransactions(r)
}

func (t *transactionMysql) Store() {
	fmt.Println("transaction_repository: store")
}

func (t *transactionMysql) Update() {
	fmt.Println("transaction_repository: update")
}

func (t *transactionMysql) Delete() {
	fmt.Println("transaction_repository: delete")
}
