package wcore

import (
	"log"
)

type Coin struct {
	ID     string
	Name   string
	Symbol string
	Price  float64
}

type CoinRepository interface {
	GetPrices(ids []string) ([]float64, error)
	GetCoins() ([]Coin, error)
}

type CoinUsecase struct {
	coinRepo CoinRepository
}

func NewCoinUsecase(c CoinRepository) *CoinUsecase {
	return &CoinUsecase{
		coinRepo: c,
	}
}

func (c *CoinUsecase) GetCoins() ([]Coin, error) {
	coins, err := c.coinRepo.GetCoins()
	if err != nil {
		log.Println(err)
		return nil, NewCoinAPIError(err)
	}

	return coins, nil
}

func (c *CoinUsecase) GetPrices(ids []string) ([]float64, error) {
	prices, err := c.coinRepo.GetPrices(ids)
	if err != nil {
		log.Println(err)
		return nil, NewCoinAPIError(err)
	}

	return prices, nil
}
