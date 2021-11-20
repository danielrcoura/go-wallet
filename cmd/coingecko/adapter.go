package coingecko

import wcore "github.com/danielrcoura/go-wallet/cmd/walletcore"

func jsonToCoins(coinsJson []coin) []wcore.Coin {
	coins := []wcore.Coin{}

	for _, c := range coinsJson {
		coins = append(coins, wcore.Coin{
			ID:     c.ID,
			Name:   c.Name,
			Symbol: c.Symbol,
		})
	}

	return coins
}

func jsonToCoinRankSummary(coinsRank []coinRankSummary) []*wcore.CoinRankSummary {
	rank := []*wcore.CoinRankSummary{}

	for _, cr := range coinsRank {
		rank = append(rank, &wcore.CoinRankSummary{
			ID:           cr.ID,
			CurrentPrice: cr.CurrentPrice,
			MarketCap:    cr.MarketCap,
			Volume:       cr.Volume,
		})
	}

	return rank
}
