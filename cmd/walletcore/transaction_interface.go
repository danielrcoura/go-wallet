package core

type TransactionRepository interface {
	FetchByWallet(walletID int) ([]Transaction, error)
	Store()
	Update()
	Delete()
}
