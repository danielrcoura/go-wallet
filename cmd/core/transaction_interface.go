package core

type TransactionRepository interface {
	Fetch()
	Store()
	Update()
	Delete()
}
