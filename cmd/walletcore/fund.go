package wcore

type FundUsecase struct {
	fundRepo      FundRepository
	walletUsecase WalletUsecase
}

func NewFundUsecase(er FundRepository, wu WalletUsecase) *FundUsecase {
	return &FundUsecase{
		fundRepo:      er,
		walletUsecase: wu,
	}
}

type MarketProp int

const (
	MarketCap MarketProp = iota
	Volume
)

type CoinRankSummary struct {
	ID           string
	CurrentPrice float64
	MarketCap    float64
	Volume       float64
}

type CoinWeight struct {
	ID     string
	Weight float64
}

type FundRepository interface {
	GetRank(size int, by MarketProp) ([]*CoinRankSummary, error)
}

func (e *FundUsecase) GetFundBalance(blacklist []string, fundSize int, by MarketProp, maxCap float64) ([]CoinWeight, error) {
	rank, err := e.getRank(blacklist, fundSize, by)
	if err != nil {
		return nil, err
	}

	coinWeights := e.getWeightsFromRank(rank, by)
	cappedIndex := 0
	cappedTotal := -1

	for cappedTotal != 0 {
		coinWeights = e.getUncappedWeights(coinWeights, cappedIndex)
		coinWeights, cappedTotal = e.applyMaxCap(coinWeights, maxCap)
		cappedIndex += cappedTotal
	}

	return coinWeights, nil
}

func (e *FundUsecase) getWeightsFromRank(rank []*CoinRankSummary, by MarketProp) []CoinWeight {
	coinWeights := []CoinWeight{}

	total := 0.0
	for _, c := range rank {
		weight := c.MarketCap
		if by == Volume {
			weight = c.Volume
		}
		coinWeights = append(coinWeights, CoinWeight{
			ID:     c.ID,
			Weight: weight,
		})
		total += weight
	}

	for i, c := range coinWeights {
		c.Weight /= total
		coinWeights[i] = c
	}

	return coinWeights
}

func (e *FundUsecase) getUncappedWeights(coinWeights []CoinWeight, cappedIndex int) []CoinWeight {
	result := coinWeights[:cappedIndex]

	uncappedPercent := 1.0
	for _, c := range coinWeights[:cappedIndex] {
		uncappedPercent -= c.Weight
	}

	uncappedTotal := 0.0
	for _, c := range coinWeights[cappedIndex:] {
		uncappedTotal += c.Weight
	}
	for _, c := range coinWeights[cappedIndex:] {
		result = append(result, CoinWeight{
			ID:     c.ID,
			Weight: c.Weight * uncappedPercent / uncappedTotal,
		})
	}

	return result
}

func (e FundUsecase) applyMaxCap(coinWeights []CoinWeight, maxCap float64) ([]CoinWeight, int) {
	result := []CoinWeight{}
	cappedTotal := 0
	for _, c := range coinWeights {
		if c.Weight > maxCap {
			cappedTotal += 1
			result = append(result, CoinWeight{
				ID:     c.ID,
				Weight: maxCap,
			})
		} else {
			result = append(result, c)
		}
	}

	return result, cappedTotal
}

func (e *FundUsecase) getRank(blacklist []string, fundSize int, by MarketProp) ([]*CoinRankSummary, error) {
	rankSize := fundSize + len(blacklist)
	rank, err := e.fundRepo.GetRank(rankSize, by)
	if err != nil {
		return nil, err
	}
	rank = e.filterBlacklist(rank, blacklist)
	rank = rank[:fundSize]

	return rank, nil
}

func (e *FundUsecase) filterBlacklist(coins []*CoinRankSummary, blacklist []string) []*CoinRankSummary {
	result := []*CoinRankSummary{}

	for _, c := range coins {
		ok := true
		for _, bl := range blacklist {
			if c.ID == bl {
				ok = false
				break
			}
		}
		if ok {
			result = append(result, c)
		}
	}

	return result
}
