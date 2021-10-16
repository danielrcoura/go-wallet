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
	ErrTransactionNotFound         = errors.New("transaction_not_found")
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

func IsDBError(err error) bool {
	_, ok := err.(dbError)
	return ok
}

type coinAPIError struct {
	err string
}

func (e coinAPIError) Error() string {
	return fmt.Sprintf("coin_api_error: %s", e.err)
}

func NewCoinAPIError(err error) coinAPIError {
	return coinAPIError{
		err: err.Error(),
	}
}

func IsCoinAPIError(err error) bool {
	_, ok := err.(coinAPIError)
	return ok
}
