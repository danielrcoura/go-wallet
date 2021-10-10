package main

import (
	http "github.com/danielrcoura/go-wallet/cmd/http"
	infra "github.com/danielrcoura/go-wallet/cmd/mysql"
	walletcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db, err := infra.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	walletMysql := infra.NewWalletMysql(db)
	walletUsecase := walletcore.NewWalletUsecase(walletMysql)

	server := http.New(walletUsecase)
	server.ListenAndServe()
}
