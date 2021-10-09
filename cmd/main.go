package main

import (
	core "github.com/danielrcoura/go-wallet/cmd/core"
	infra "github.com/danielrcoura/go-wallet/cmd/mysql"
)

func main() {
	walletMysql := infra.NewWalletMysql()
	walletUsecase := core.NewWalletUsecase(walletMysql)

	walletUsecase.Store("cripto")

	transactionMysql := infra.NewTransactionMysql()
	transactionUsecase := core.NewTransactionUsecase(transactionMysql)

	transactionUsecase.Store()
}
