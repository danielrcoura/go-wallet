package wcore

import (
	"log"
	"strings"
)

type SimpleWallet struct {
	ID   int
	Name string
}

type WalletRepository interface {
	Fetch() ([]SimpleWallet, error)
	FetchByID(id int) (*SimpleWallet, error)
	FetchByName(name string) (*SimpleWallet, error)
	Store(name string) error
	Update(id int, w SimpleWallet) error
	Delete(id int) error
}

type SimpleWalletUsecase struct {
	walletRepo WalletRepository
}

func NewSimpleWalletUsecase(w WalletRepository) *SimpleWalletUsecase {
	return &SimpleWalletUsecase{
		walletRepo: w,
	}
}

func (wl *SimpleWalletUsecase) Fetch() ([]SimpleWallet, error) {
	wallets, err := wl.walletRepo.Fetch()
	if err != nil {
		return nil, NewDBError(err)
	}

	return wallets, nil
}

func (wl *SimpleWalletUsecase) FetchByID(id int) (*SimpleWallet, error) {
	w, err := wl.walletRepo.FetchByID(id)
	if err != nil {
		log.Println(err)
		return nil, NewDBError(err)
	} else if w == nil {
		log.Println(ErrWalletNotFound)
		return nil, ErrWalletNotFound
	}

	return w, nil
}

func (wl *SimpleWalletUsecase) Store(name string) error {
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

func (wl *SimpleWalletUsecase) Update(id int, w SimpleWallet) error {
	name, err := wl.handleName(w.Name)
	if err != nil {
		log.Println(err)
		return err
	}
	w.Name = name

	_, err = wl.FetchByID(id)
	if err != nil {
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

	if err = wl.walletRepo.Update(id, w); err != nil {
		return NewDBError(err)
	}

	return nil
}

func (wl *SimpleWalletUsecase) Delete(id int) error {
	_, err := wl.FetchByID(id)
	if err != nil {
		return err
	}

	if err := wl.walletRepo.Delete(id); err != nil {
		log.Println(err)
		return NewDBError(err)
	}

	return nil
}

func (wl *SimpleWalletUsecase) handleName(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", ErrInvalidWalletName
	}

	return name, nil
}

func (wl *SimpleWalletUsecase) checkWalletExists(name string) (bool, error) {
	w, err := wl.walletRepo.FetchByName(name)
	if err != nil {
		return false, NewDBError(err)
	}
	return w != nil, nil
}
