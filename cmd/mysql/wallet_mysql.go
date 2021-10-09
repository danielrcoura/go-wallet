package infra

import (
	"fmt"

	"github.com/danielrcoura/go-wallet/cmd/core"
)

type walletMysql struct{}

func NewWalletMysql() *walletMysql {
	return &walletMysql{}
}

func (wl *walletMysql) Store(name string) error {
	fmt.Println("repository: store")
	return nil
}

func (wl *walletMysql) Update(id int64, w core.Wallet) error {
	fmt.Println("repository: update")
	return nil
}

func (wl *walletMysql) Delete(id int64) error {
	fmt.Println("repository: delete")
	return nil
}

func (wl *walletMysql) Fetch() ([]core.Wallet, error) {
	fmt.Println("repository: fetch")
	return nil, nil
}
