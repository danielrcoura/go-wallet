package mysql

import (
	"database/sql"
	"fmt"
	"strings"

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

func (t *transactionMysql) Update(id int, transaction wcore.Transaction) error {
	fields := []string{}
	data := []interface{}{}

	if transaction.Ticker != "" {
		data = append(data, transaction.Ticker)
		fields = append(fields, "ticker")
	}
	if transaction.Operation.Check() {
		data = append(data, transaction.Operation)
		fields = append(fields, "operation")
	}
	if transaction.Quantity > 0 {
		data = append(data, transaction.Quantity)
		fields = append(fields, "quantity")
	}
	if transaction.Price > 0 {
		data = append(data, transaction.Price)
		fields = append(fields, "price")
	}
	if transaction.Date != nil {
		data = append(data, transaction.Date)
		fields = append(fields, "date")
	}

	for i, f := range fields {
		fields[i] = f + "=?"
	}
	fieldQuery := strings.Join(fields, ",")

	query := fmt.Sprintf(`UPDATE transactions 
						  SET %s
						  WHERE id=?`, fieldQuery)
	data = append(data, id)

	fmt.Println(query, data)
	_, err := t.db.Exec(query, data...)
	if err != nil {
		return err
	}

	return nil
}

func (t *transactionMysql) Delete(id int) error {
	fmt.Println("transaction_repository: delete")
	return nil
}
