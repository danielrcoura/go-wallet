package http

import (
	walletcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
)

func walletReqToWallet(wReq walletReq) walletcore.Wallet {
	return walletcore.Wallet{
		Name: wReq.Name,
	}
}
