package infra

import (
	"fmt"
)

type transactionMysql struct{}

func NewTransactionMysql() *transactionMysql {
	return &transactionMysql{}
}

func (w *transactionMysql) Store() {
	fmt.Println("transaction_repository: store")
}

func (w *transactionMysql) Update() {
	fmt.Println("transaction_repository: update")
}

func (w *transactionMysql) Delete() {
	fmt.Println("transaction_repository: delete")
}

func (w *transactionMysql) Fetch() {
	fmt.Println("transaction_repository: fetch")
}
