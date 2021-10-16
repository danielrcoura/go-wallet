package main

import (
	"log"

	"github.com/danielrcoura/go-wallet/cmd/coingecko"
	http "github.com/danielrcoura/go-wallet/cmd/http"
	mysql "github.com/danielrcoura/go-wallet/cmd/mysql"
	wcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	db, err := mysql.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	walletMysql := mysql.NewWalletMysql(db)
	walletUsecase := wcore.NewWalletUsecase(walletMysql)

	transactionMysql := mysql.NewTransactionMysql(db)
	transactionUsecase := wcore.NewTransactionUsecase(transactionMysql, walletUsecase)

	coinGecko := coingecko.New()
	coinUsecase := wcore.NewCoinUsecase(coinGecko)

	log.Println("Starting server...")
	server := http.New(
		walletUsecase,
		transactionUsecase,
		coinUsecase,
	)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
