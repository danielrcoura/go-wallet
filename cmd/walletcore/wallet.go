package wcore

import "fmt"

type CoinSummary struct {
	TotalQuantity float64
	CurrentPrice  float64
	AvgPrice      float64
}

func (cs CoinSummary) Value() float64 {
	return cs.CurrentPrice * cs.TotalQuantity
}

type Wallet struct {
	Wallet  SimpleWallet
	Summary map[string]CoinSummary
}

func (w Wallet) Value() float64 {
	value := 0.0
	for _, c := range w.Summary {
		value += c.Value()
	}
	return value
}

func (w Wallet) AvgPrice() float64 {
	price := 0.0
	for _, c := range w.Summary {
		price += c.AvgPrice
	}
	return price
}

type WalletUsecase struct {
	transactionUsecase TransactionUsecase
	swalletUsecase     SimpleWalletUsecase
	coinUsecase        CoinUsecase
}

func NewWalletUsecase(tu TransactionUsecase, sw SimpleWalletUsecase, c CoinUsecase) *WalletUsecase {
	return &WalletUsecase{
		transactionUsecase: tu,
		swalletUsecase:     sw,
		coinUsecase:        c,
	}
}

func (wu *WalletUsecase) BuildWallet(id int) (*Wallet, error) {
	swallet, err := wu.swalletUsecase.FetchByID(id)
	if err != nil {
		return nil, err
	}

	sum, err := wu.summariseWalletTransactions(swallet.ID)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		Wallet:  *swallet,
		Summary: sum,
	}, nil
}

func (wu *WalletUsecase) BuildWallets() ([]*Wallet, error) {
	swallets, err := wu.swalletUsecase.Fetch()
	if err != nil {
		return nil, err
	}

	wallets := []*Wallet{}

	for _, sw := range swallets {
		sum, err := wu.summariseWalletTransactions(sw.ID)
		if err != nil {
			return nil, err
		}

		w := Wallet{
			Wallet:  sw,
			Summary: sum,
		}

		wallets = append(wallets, &w)
	}

	return wallets, nil
}

func (wu *WalletUsecase) Rebalance(walletId int, goal []CoinWeight, supply float64) (map[string]float64, error) {
	result := map[string]float64{}

	wallet, err := wu.BuildWallet(walletId)
	if err != nil {
		return nil, err
	}
	currBalance := wu.getBalance(wallet.Summary)

	for _, c := range currBalance {
		fmt.Printf("%v %f\n", c.ID, c.Weight)
	}
	fmt.Println("---------")

	// get caixa
	idealSupply := 0.0
	for _, g := range goal {
		s := wallet.Summary[g.ID].Value() / g.Weight
		if s > idealSupply {
			idealSupply = s
		}
	}
	idealSupply -= wallet.Value()

	fmt.Println("idealSupply: ", idealSupply)
	fmt.Println("----------")

	// sell
	for _, c := range currBalance {
		found := false
		for _, g := range goal {
			if c.ID == g.ID {
				found = true
			}
		}
		if !found {
			result[c.ID] = -wallet.Summary[c.ID].Value()
		}
	}

	// balance with actual supply
	totalSupply := wallet.Value() + supply
	for _, g := range goal {
		coin, ok := wallet.Summary[g.ID]
		result[g.ID] = totalSupply * g.Weight
		if ok {
			result[g.ID] -= coin.Value()
		}
	}

	return result, nil
}

func (wu *WalletUsecase) getBalance(wallet map[string]CoinSummary) []CoinWeight {
	result := []CoinWeight{}
	total := 0.0

	for _, c := range wallet {
		total += c.Value()
	}

	for id, c := range wallet {
		result = append(result, CoinWeight{
			ID:     id,
			Weight: c.Value() / total,
		})
	}

	return result
}

func (wu *WalletUsecase) summariseWalletTransactions(walletId int) (map[string]CoinSummary, error) {
	transactionsByCoin, err := wu.groupTransactionsByCoin(walletId)
	if err != nil {
		return nil, err
	}

	summary := map[string]CoinSummary{}
	for coinId, transactions := range transactionsByCoin {
		summary[coinId] = CoinSummary{
			AvgPrice:      wu.calculateAvgPrice(transactions),
			TotalQuantity: wu.calculateQuantity(transactions),
		}
	}

	summary, err = wu.fillCurrentPrices(summary)
	if err != nil {
		return nil, err
	}

	return summary, nil
}

func (wu *WalletUsecase) calculateQuantity(transactions []Transaction) float64 {
	qtd := 0.0
	for _, t := range transactions {
		qtd += t.Quantity
	}
	return qtd
}

func (wu *WalletUsecase) calculateAvgPrice(transactions []Transaction) float64 {
	buyPrice := 0.0
	quantityBuy := 0.0

	for _, t := range transactions {
		if t.Quantity > 0 {
			buyPrice += t.Quantity * t.Price
			quantityBuy += t.Quantity
		}
	}

	return buyPrice / quantityBuy
}

func (wu *WalletUsecase) groupTransactionsByCoin(walletId int) (map[string][]Transaction, error) {
	transactions, err := wu.transactionUsecase.FetchByWallet(walletId)
	if err != nil {
		return nil, err
	}

	transactionsByCoin := map[string][]Transaction{}
	for _, t := range transactions {
		tc := transactionsByCoin[t.Ticker]
		transactionsByCoin[t.Ticker] = append(tc, t)
	}

	return transactionsByCoin, nil
}

func (wu *WalletUsecase) fillCurrentPrices(summary map[string]CoinSummary) (map[string]CoinSummary, error) {
	tickers := []string{}
	for ticker := range summary {
		tickers = append(tickers, ticker)
	}

	prices, err := wu.coinUsecase.GetPrices(tickers)
	if err != nil {
		return nil, err
	}

	for i, t := range tickers {
		s := summary[t]
		s.CurrentPrice = prices[i]
		summary[t] = s
	}

	return summary, nil
}
