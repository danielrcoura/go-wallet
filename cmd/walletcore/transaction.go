package wcore

import (
	"log"
	"strings"
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

func (op Operation) Check() bool {
	return op == Buy || op == Sell
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
	Date      *time.Time
}

type TransactionRepository interface {
	FetchByWallet(walletID int) ([]Transaction, error)
	Store(walletID int, transaction Transaction) error
	Update(id int, transaction Transaction) error
	Delete(id int) error
}

type TransactionUsecase struct {
	transactionRepo TransactionRepository
	walletUsecase   *WalletUsecase
}

func NewTransactionUsecase(t TransactionRepository, w *WalletUsecase) *TransactionUsecase {
	return &TransactionUsecase{
		transactionRepo: t,
		walletUsecase:   w,
	}
}

func (tu *TransactionUsecase) FetchByWallet(walletID int) ([]Transaction, error) {
	_, err := tu.walletUsecase.FetchByID(walletID)
	if err != nil {
		return nil, err
	}

	transactions, err := tu.transactionRepo.FetchByWallet(walletID)
	if err != nil {
		log.Print(err)
		return nil, NewDBError(err)
	}

	return transactions, nil
}

func (tu *TransactionUsecase) Store(walletID int, transaction Transaction) error {
	_, err := tu.walletUsecase.FetchByID(walletID)
	if err != nil {
		return err
	}

	transaction, err = tu.checkTransaction(transaction)
	if err != nil {
		log.Print(err)
		return err
	}

	if err := tu.transactionRepo.Store(walletID, transaction); err != nil {
		log.Print(err)
		return NewDBError(err)
	}

	return nil
}

func (tu *TransactionUsecase) Update(id int, transaction Transaction) error {
	// TODO: check if transaction exists

	if err := tu.transactionRepo.Update(id, transaction); err != nil {
		log.Println(err)
		return NewDBError(err)
	}

	return nil
}

func (tu *TransactionUsecase) Delete(id int) error {
	if err := tu.transactionRepo.Delete(id); err != nil {
		return NewDBError(err)
	}

	return nil
}

func (tu *TransactionUsecase) checkTransaction(t Transaction) (Transaction, error) {
	t.Ticker = strings.TrimSpace(t.Ticker)

	if t.Ticker == "" {
		return t, ErrInvalidTransactionTicker
	}

	if !t.Operation.Check() {
		return t, ErrInvalidTransactionOperation
	}

	if t.Quantity <= 0 {
		return t, ErrInvalidTransactionQuantity
	}

	if t.Price <= 0 {
		return t, ErrInvalidTransactionPrice
	}

	if !t.Date.Before(time.Now()) && !t.Date.Equal(time.Now()) {
		return t, ErrInvalidTransactionDate
	}

	return t, nil
}
