package main

import (
	"fmt"
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
	simpleWalletUsecase := wcore.NewSimpleWalletUsecase(walletMysql)

	transactionMysql := mysql.NewTransactionMysql(db)
	transactionUsecase := wcore.NewTransactionUsecase(transactionMysql, simpleWalletUsecase)

	coinGecko := coingecko.New()
	coinUsecase := wcore.NewCoinUsecase(coinGecko)

	walletUsecase := wcore.NewWalletUsecase(*transactionUsecase, *simpleWalletUsecase, *coinUsecase)

	etfUsecase := wcore.NewFundUsecase(coinGecko, *walletUsecase)

	blacklist := []string{
		"tether",
		"binance-usd",
		"usd-coin",
		"dai",
		"bitcoin",
		"ethereum",
		"ripple",
		"wrapped-bitcoin",
		"internet-computer",
	}
	goal, err := etfUsecase.GetFundBalance(blacklist, 20, wcore.MarketCap, 0.1)
	if err != nil {
		log.Fatal(err)
	}

	r, err := walletUsecase.Rebalance(2, goal, 1000)
	if err != nil {
		log.Fatal(err)
	}

	for c, v := range r {
		fmt.Printf("%v %f\n", c, v)
	}

	log.Println("Starting server...")
	server := http.New(
		simpleWalletUsecase,
		transactionUsecase,
		coinUsecase,
		walletUsecase,
	)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
