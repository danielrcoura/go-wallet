package main

import (
	"log"

	http "github.com/danielrcoura/go-wallet/cmd/http"
	infra "github.com/danielrcoura/go-wallet/cmd/mysql"
	walletcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := infra.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	walletMysql := infra.NewWalletMysql(db)
	walletUsecase := walletcore.NewWalletUsecase(walletMysql)

	server := http.New(walletUsecase)
	log.Println("Starting server...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
