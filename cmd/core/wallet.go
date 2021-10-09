package core

import (
	"fmt"
	"strings"
)

type Wallet struct {
	ID           int64
	Name         string
	Transactions []Transaction
}

type walletUsecase struct {
	walletRepo WalletRepository
}

func NewWalletUsecase(w WalletRepository) *walletUsecase {
	return &walletUsecase{
		walletRepo: w,
	}
}

func (wl *walletUsecase) Store(name string) error {
	name, err := handleName(name)
	if err != nil {
		return err
	}

	if err := wl.walletRepo.Store(name); err != nil {
		return err
	}

	return nil
}

func (wl *walletUsecase) Update(id int64, w Wallet) error {
	name, err := handleName(w.Name)
	if err != nil {
		return err
	}
	w.Name = name

	if err = wl.walletRepo.Update(id, w); err != nil {
		return err
	}

	return nil
}

func (wl *walletUsecase) Delete(id int64) error {
	if err := wl.walletRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (wl *walletUsecase) Fetch() ([]Wallet, error) {
	wallets, err := wl.walletRepo.Fetch()
	if err != nil {
		return nil, err
	}

	return wallets, nil
}

func handleName(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", fmt.Errorf("'name' cannot be null")
	}

	return name, nil
}
