package coingecko

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	wcore "github.com/danielrcoura/go-wallet/cmd/walletcore"
)

const (
	BASE_URL = "https://api.coingecko.com/api/v3"
	CURRENCY = "usd"
)

type coin struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type coinGecko struct {
}

func New() coinGecko {
	return coinGecko{}
}

func (cg coinGecko) GetCoins() ([]wcore.Coin, error) {
	url := fmt.Sprintf("%s/coins/list", BASE_URL)

	var coinsJson []coin
	err := makeRequest(url, &coinsJson)
	if err != nil {
		return nil, err
	}

	return jsonToCoins(coinsJson), nil
}

func (cg coinGecko) GetPrices(ids []string) ([]float64, error) {
	url := fmt.Sprintf(
		"%s/simple/price?ids=%s&vs_currencies=%s",
		BASE_URL,
		strings.Join(ids, ","),
		CURRENCY,
	)

	var json map[string]map[string]float64
	err := makeRequest(url, &json)
	if err != nil {
		return nil, err
	}

	result := []float64{}
	for _, id := range ids {
		v, ok := json[id]
		if ok {
			result = append(result, v[CURRENCY])
		} else {
			result = append(result, -1)
		}
	}

	return result, nil
}

func makeRequest(url string, target interface{}) error {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	r, err := netClient.Get(url)
	if err != nil {
		return err
	}
	r.Close = true
	defer r.Body.Close()

	// A 429 error is a rate limit error.
	// CoinGecko limits requests to 10 calls per second per IP address.
	if r.StatusCode < 200 || r.StatusCode >= 300 {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("status: %v, error: %v", r.StatusCode, string(body))
	}

	return json.NewDecoder(r.Body).Decode(target)
}
