package wcore

type CoinSummary struct {
	TotalQuantity float64
	CurrentPrice  float64
	AvgPrice      float64
}

type Wallet struct {
	Wallet  SimpleWallet
	Summary map[string]CoinSummary
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
