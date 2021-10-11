package infra

import (
	"database/sql"
	"fmt"

	adapter "github.com/danielrcoura/go-wallet/cmd/adapters"
	walletcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
	"github.com/danielrcoura/go-wallet/cmd/walleterror"
)

type walletMysql struct {
	db *sql.DB
}

func NewWalletMysql(db *sql.DB) *walletMysql {
	return &walletMysql{
		db: db,
	}
}

func (wl *walletMysql) Fetch() ([]walletcore.Wallet, error) {
	r, err := wl.db.Query("SELECT * FROM wallets")
	if err != nil {
		return nil, walleterror.NewDBError(err)
	}

	wallets, err := adapter.RowsToWallets(r)
	if err != nil {
		return nil, walleterror.NewDBError(err)
	}

	return wallets, nil
}

func (wl *walletMysql) FetchByID(id string) (*walletcore.Wallet, error) {
	r, err := wl.db.Query("SELECT * FROM wallets WHERE id=?", id)
	if err != nil {
		return nil, walleterror.NewDBError(err)
	}

	wallets, err := adapter.RowsToWallets(r)
	if err != nil {
		return nil, walleterror.NewDBError(err)
	}

	if len(wallets) < 1 {
		return nil, walleterror.ErrWalletDoesNotExists
	}

	return &wallets[0], nil
}

func (wl *walletMysql) FetchByName(name string) (*walletcore.Wallet, error) {
	r, err := wl.db.Query("SELECT * FROM wallets WHERE name=?", name)
	if err != nil {
		return nil, walleterror.NewDBError(err)
	}

	wallets, err := adapter.RowsToWallets(r)
	if err != nil {
		return nil, walleterror.NewDBError(err)
	}

	if len(wallets) < 1 {
		return nil, walleterror.ErrWalletDoesNotExists
	}

	return &wallets[0], nil
}

func (wl *walletMysql) Store(name string) error {
	stmt, err := wl.db.Prepare("INSERT INTO wallets (name) VALUES (?)")
	defer stmt.Close()
	if err != nil {
		return walleterror.NewDBError(err)
	}

	_, err = stmt.Exec(name)
	if err != nil {
		return walleterror.NewDBError(err)
	}

	return nil
}

func (wl *walletMysql) Update(id int, w walletcore.Wallet) error {
	stmt, err := wl.db.Prepare("UPDATE wallets SET name=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return walleterror.NewDBError(err)
	}

	_, err = stmt.Exec(w.Name, id)
	if err != nil {
		return walleterror.NewDBError(err)
	}

	return nil
}

func (wl *walletMysql) Delete(id int) error {
	fmt.Println("repository: delete")
	return nil
}
