package mysql

import (
	"database/sql"

	wcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
)

type walletMysql struct {
	db *sql.DB
}

func NewWalletMysql(db *sql.DB) *walletMysql {
	return &walletMysql{
		db: db,
	}
}

func (wl *walletMysql) Fetch() ([]wcore.Wallet, error) {
	r, err := wl.db.Query("SELECT * FROM wallets")
	if err != nil {
		return nil, err
	}

	wallets, err := RowsToWallets(r)
	if err != nil {
		return nil, err
	}

	return wallets, nil
}

func (wl *walletMysql) FetchByID(id int) (*wcore.Wallet, error) {
	r, err := wl.db.Query("SELECT * FROM wallets WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	wallets, err := RowsToWallets(r)
	if err != nil {
		return nil, err
	}

	if len(wallets) == 0 {
		return nil, nil
	}

	return &wallets[0], nil
}

func (wl *walletMysql) FetchByName(name string) (*wcore.Wallet, error) {
	r, err := wl.db.Query("SELECT * FROM wallets WHERE name=?", name)
	if err != nil {
		return nil, err
	}

	wallets, err := RowsToWallets(r)
	if err != nil {
		return nil, err
	}

	if len(wallets) == 0 {
		return nil, nil
	}

	return &wallets[0], nil
}

func (wl *walletMysql) Store(name string) error {
	_, err := wl.db.Exec("INSERT INTO wallets (name) VALUES (?)", name)
	if err != nil {
		return err
	}

	return nil
}

func (wl *walletMysql) Update(id int, w wcore.Wallet) error {
	_, err := wl.db.Exec("UPDATE wallets SET name=? WHERE id=?", w.Name, id)
	if err != nil {
		return err
	}

	return nil
}

func (wl *walletMysql) Delete(id int) error {
	_, err := wl.db.Exec("DELETE FROM wallets WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}
