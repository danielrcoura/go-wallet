package http

import (
	wcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
)

func walletReqToWallet(wReq walletReq) wcore.Wallet {
	return wcore.Wallet{
		Name: wReq.Name,
	}
}
