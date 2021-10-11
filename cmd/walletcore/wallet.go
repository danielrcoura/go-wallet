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
	FetchByID(id string) (*Wallet, error)
	FetchByName(name string) (*Wallet, error)
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

func (wl *WalletUsecase) Fetch() ([]Wallet, error) {
	return wl.walletRepo.Fetch()
}

func (wl *WalletUsecase) FetchByID(id string) (*Wallet, error) {
	return wl.walletRepo.FetchByID(id)
}

func (wl *WalletUsecase) Store(name string) error {
	log.Printf("Creating wallet %s...\n", name)
	name, err := wl.handleName(name)
	if err != nil {
		log.Println(err)
		return err
	}

	exists, err := wl.checkWalletExists(name)
	if err != nil {
		log.Println(err)
		return err
	} else if exists {
		log.Println(walleterror.ErrWalletAlreadyExists)
		return walleterror.ErrWalletAlreadyExists
	}

	if err := wl.walletRepo.Store(name); err != nil {
		log.Println(err)
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

func (wl *WalletUsecase) handleName(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", walleterror.ErrInvalidWalletName
	}

	return name, nil
}

func (wl *WalletUsecase) checkWalletExists(name string) (bool, error) {
	w, err := wl.walletRepo.FetchByName(name)
	if err != nil && walleterror.IsdbError(err) {
		return false, err
	}
	return w != nil, nil
}
