package mysql

import (
	"database/sql"

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

	wallets, err := RowsToWallets(r)
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

	wallets, err := RowsToWallets(r)
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

	wallets, err := RowsToWallets(r)
	if err != nil {
		return nil, walleterror.NewDBError(err)
	}

	if len(wallets) < 1 {
		return nil, walleterror.ErrWalletDoesNotExists
	}

	return &wallets[0], nil
}

func (wl *walletMysql) Store(name string) error {
	_, err := wl.db.Exec("INSERT INTO wallets (name) VALUES (?)", name)
	if err != nil {
		return walleterror.NewDBError(err)
	}

	return nil
}

func (wl *walletMysql) Update(id int, w walletcore.Wallet) error {
	_, err := wl.db.Exec("UPDATE wallets SET name=? WHERE id=?", w.Name, id)
	if err != nil {
		return walleterror.NewDBError(err)
	}

	return nil
}

func (wl *walletMysql) Delete(id int) error {
	_, err := wl.db.Exec("DELETE FROM wallets WHERE id=?", id)
	if err != nil {
		return walleterror.NewDBError(err)
	}

	return nil
}
