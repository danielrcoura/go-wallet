package core

import (
	"log"
	"strings"

	walleterror "github.com/danielrcoura/go-wallet/cmd/walleterror"
)

type Wallet struct {
	ID           int64
	Name         string
	Transactions []Transaction
}

type WalletRepository interface {
	Fetch() ([]Wallet, error)
	Store(name string) error
	Update(id int64, w Wallet) error
	Delete(id int64) error
}

type WalletUsecase struct {
	walletRepo WalletRepository
}

func NewWalletUsecase(w WalletRepository) *WalletUsecase {
	return &WalletUsecase{
		walletRepo: w,
	}
}

func (wl *WalletUsecase) Store(name string) error {
	// TODO: check if there's another wallet with the same name
	log.Printf("Creating wallet %s...\n", name)
	name, err := wl.handleName(name)
	if err != nil {
		log.Println(err)
		return err
	}

	if err := wl.walletRepo.Store(name); err != nil {
		return err
	}

	return nil
}

func (wl *WalletUsecase) Update(id int64, w Wallet) error {
	name, err := wl.handleName(w.Name)
	if err != nil {
		return err
	}
	w.Name = name

	if err = wl.walletRepo.Update(id, w); err != nil {
		return err
	}

	return nil
}

func (wl *WalletUsecase) Delete(id int64) error {
	if err := wl.walletRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (wl *WalletUsecase) Fetch() ([]Wallet, error) {
	return wl.walletRepo.Fetch()
}

func (wl *WalletUsecase) handleName(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", walleterror.ErrInvalidWalletName
	}

	return name, nil
}
