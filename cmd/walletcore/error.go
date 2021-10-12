package wcore

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidWalletName   = errors.New("invalid_wallet_name")
	ErrWalletAlreadyExists = errors.New("wallet_already_exists")
	ErrWalletNotFound      = errors.New("wallet_not_found")

	ErrInvalidTransactionTicker    = errors.New("invalid_transaction_ticker")
	ErrInvalidTransactionOperation = errors.New("invalid_transaction_operation")
	ErrInvalidTransactionQuantity  = errors.New("invalid_transaction_quantity")
	ErrInvalidTransactionPrice     = errors.New("invalid_transaction_price")
	ErrInvalidTransactionDate      = errors.New("invalid_transaction_date")
)

type dbError struct {
	err string
}

func (e dbError) Error() string {
	return fmt.Sprintf("database_error: %s", e.err)
}

func NewDBError(err error) dbError {
	return dbError{
		err: err.Error(),
	}
}

func IsdbError(err error) bool {
	_, ok := err.(dbError)
	return ok
}
