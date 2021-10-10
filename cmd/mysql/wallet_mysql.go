package infra

import (
	"database/sql"
	"fmt"

	adapter "github.com/danielrcoura/go-wallet/cmd/adapters"
	walletcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
)

type walletMysql struct {
	db *sql.DB
}

func NewWalletMysql(db *sql.DB) *walletMysql {
	return &walletMysql{
		db: db,
	}
}

func (wl *walletMysql) Store(name string) error {
	stmt, err := wl.db.Prepare("INSERT INTO wallets (name) VALUES (?)")
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name)
	if err != nil {
		return err
	}

	return nil
}

func (wl *walletMysql) Update(id int64, w walletcore.Wallet) error {
	fmt.Println("repository: update")
	return nil
}

func (wl *walletMysql) Delete(id int64) error {
	fmt.Println("repository: delete")
	return nil
}

func (wl *walletMysql) Fetch() ([]walletcore.Wallet, error) {
	r, err := wl.db.Query("SELECT * FROM wallets")
	if err != nil {
		return nil, err
	}

	wallets, err := adapter.RowsToWallets(r)
	if err != nil {
		return nil, err
	}

	return wallets, nil
}
