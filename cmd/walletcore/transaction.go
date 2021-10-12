package wcore

import (
	"fmt"
	"log"
	"time"
)

type Operation int

const (
	Buy Operation = iota + 1
	Sell
)

func (op Operation) String() string {
	switch op {
	case Buy:
		return "buy"
	case Sell:
		return "sell"
	default:
		return ""
	}
}

func StringToOperation(op string) Operation {
	switch op {
	case Buy.String():
		return Buy
	case Sell.String():
		return Sell
	default:
		return 0
	}
}

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
	Store(walletID int, transaction Transaction) error
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

func (tu *TransactionUsecase) FetchByWallet(walletID int) ([]Transaction, error) {
	// TODO: wallet not found
	transactions, err := tu.transactionRepo.FetchByWallet(walletID)
	if err != nil {
		log.Print(err)
		return nil, NewDBError(err)
	}

	return transactions, nil
}

func (tu *TransactionUsecase) Store(walletID int, transaction Transaction) error {
	// TODO: validate
	if err := tu.transactionRepo.Store(walletID, transaction); err != nil {
		log.Print(err)
		return NewDBError(err)
	}

	return nil
}

func (tu *TransactionUsecase) Update() {
	fmt.Println("transaction_usecase: update")
}

func (tu *TransactionUsecase) Delete() {
	fmt.Println("transaction_usecase: delete")
}
