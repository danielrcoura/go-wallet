package wcore

type CoinSummary struct {
	TotalQuantity float64
	AvgPrice      float64
}

type Wallet struct {
	Wallet  SimpleWallet
	Summary map[string]*CoinSummary
}

type WalletUsecase struct {
	transactionUsecase TransactionUsecase
	swalletUsecase     SimpleWalletUsecase
}

func NewWalletUsecase(tu TransactionUsecase, sw SimpleWalletUsecase) *WalletUsecase {
	return &WalletUsecase{
		transactionUsecase: tu,
		swalletUsecase:     sw,
	}
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

func (wu *WalletUsecase) summariseWalletTransactions(walletId int) (map[string]*CoinSummary, error) {
	transactions, err := wu.transactionUsecase.FetchByWallet(walletId)
	if err != nil {
		return nil, err
	}

	summary := map[string]*CoinSummary{}

	for _, t := range transactions {
		ts, ok := summary[t.Ticker]
		if ok {
			ts.TotalQuantity += t.Quantity
		} else {
			summary[t.Ticker] = &CoinSummary{
				TotalQuantity: t.Quantity,
			}
		}
	}

	return summary, nil
}
