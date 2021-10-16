package test

import (
	"testing"

	"github.com/danielrcoura/go-wallet/cmd/coingecko"
	wcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
)

func TestGetPrices(t *testing.T) {
	cg := coingecko.New()

	c := wcore.NewCoinUsecase(cg)

	coins, err := c.GetCoins()
	if err != nil {
		t.Error("Error: ", err)
	}

	ids := []string{}
	for _, c := range coins[:5] {
		ids = append(ids, c.ID)
	}

	prices, err := c.GetPrices(ids)
	if err != nil {
		t.Error("Error: ", err)
	}

	if len(prices) != len(ids) {
		t.Errorf("IDs: %v, Prices: %v", ids, prices)
	}

	for i, p := range prices {
		if p < 0 {
			t.Errorf("%v price not found", ids[i])
		}
	}
}
