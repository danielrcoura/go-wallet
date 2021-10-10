package core

import (
	"fmt"
	"time"
)

type Operation int

const (
	Buy Operation = iota + 1
	Sell
)

type Transaction struct {
	ID        int64
	Ticker    string
	Operation Operation
	Quantity  float64
	Price     float64
	Date      time.Time
}

type transactionUsecase struct {
	transactionRepo TransactionRepository
}

func NewTransactionUsecase(t TransactionRepository) *transactionUsecase {
	return &transactionUsecase{
		transactionRepo: t,
	}
}

func (t *transactionUsecase) FetchByWallet(walletID int) ([]Transaction, error) {
	return t.transactionRepo.FetchByWallet(walletID)
}

func (t *transactionUsecase) Store() {
	fmt.Println("transaction_usecase: store")
}

func (t *transactionUsecase) Update() {
	fmt.Println("transaction_usecase: update")
}

func (t *transactionUsecase) Delete() {
	fmt.Println("transaction_usecase: delete")
}
