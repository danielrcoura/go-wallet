package core

type WalletRepository interface {
	Fetch() ([]Wallet, error)
	Store(name string) error
	Update(id int64, w Wallet) error
	Delete(id int64) error
}
