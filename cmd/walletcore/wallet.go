package wcore

import (
	"log"
	"strings"
)

type Wallet struct {
	ID   int
	Name string
}

type WalletRepository interface {
	Fetch() ([]Wallet, error)
	FetchByID(id string) (*Wallet, error)
	FetchByName(name string) (*Wallet, error)
	Store(name string) error
	Update(id int, w Wallet) error
	Delete(id int) error
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
	wallets, err := wl.walletRepo.Fetch()
	if err != nil {
		return nil, NewDBError(err)
	}

	return wallets, nil
}

func (wl *WalletUsecase) FetchByID(id string) (*Wallet, error) {
	w, err := wl.walletRepo.FetchByID(id)
	if err != nil {
		return nil, NewDBError(err)
	}

	return w, nil
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
		log.Println(ErrWalletAlreadyExists)
		return ErrWalletAlreadyExists
	}

	if err := wl.walletRepo.Store(name); err != nil {
		log.Println(err)
		return NewDBError(err)
	}

	return nil
}

func (wl *WalletUsecase) Update(id int, w Wallet) error {
	name, err := wl.handleName(w.Name)
	if err != nil {
		log.Println(err)
		return err
	}
	w.Name = name

	exists, err := wl.checkWalletExists(name)
	if err != nil {
		log.Println(err)
		return err
	} else if exists {
		log.Println(ErrWalletAlreadyExists)
		return ErrWalletAlreadyExists
	}

	// TODO: wallet not found

	if err = wl.walletRepo.Update(id, w); err != nil {
		return NewDBError(err)
	}

	return nil
}

func (wl *WalletUsecase) Delete(id int) error {
	// TODO: wallet not found

	if err := wl.walletRepo.Delete(id); err != nil {
		log.Println(err)
		return NewDBError(err)
	}

	return nil
}

func (wl *WalletUsecase) handleName(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", ErrInvalidWalletName
	}

	return name, nil
}

func (wl *WalletUsecase) checkWalletExists(name string) (bool, error) {
	w, err := wl.walletRepo.FetchByName(name)
	if err != nil {
		return false, NewDBError(err)
	}
	return w != nil, nil
}
