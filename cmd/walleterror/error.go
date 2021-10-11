package walleterror

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidWalletName   = errors.New("invalid_wallet_name")
	ErrWalletAlreadyExists = errors.New("wallet_already_exists")
	ErrWalletDoesNotExists = errors.New("wallet_does_not_exists")
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
