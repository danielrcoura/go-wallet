package walleterror

import "errors"

var (
	ErrInvalidWalletName   = errors.New("invalid_wallet_name")
	ErrWalletAlreadyExists = errors.New("wallet_already_exists")
)
