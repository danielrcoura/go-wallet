package wcore

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
	ID        int
	Ticker    string
	Operation Operation
	Quantity  float64
	Price     float64
	Date      time.Time
}

type TransactionRepository interface {
	FetchByWallet(walletID int) ([]Transaction, error)
	Store()
	Update()
	Delete()
}

type TransactionUsecase struct {
	transactionRepo TransactionRepository
}

func NewTransactionUsecase(t TransactionRepository) *TransactionUsecase {
	return &TransactionUsecase{
		transactionRepo: t,
	}
}

func (t *TransactionUsecase) FetchByWallet(walletID int) ([]Transaction, error) {
	return t.transactionRepo.FetchByWallet(walletID)
}

func (t *TransactionUsecase) Store() {
	fmt.Println("transaction_usecase: store")
}

func (t *TransactionUsecase) Update() {
	fmt.Println("transaction_usecase: update")
}

func (t *TransactionUsecase) Delete() {
	fmt.Println("transaction_usecase: delete")
}
