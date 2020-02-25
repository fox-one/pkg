package f1exapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadAssetTickers(t *testing.T) {
	tickers, err := ReadAssetTickers(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

	if assert.NotEmpty(t, tickers) {
		for _, ticker := range tickers {
			t.Log(ticker.AssetID)
			t.Log(ticker.Symbol)
			t.Log(ticker.Price, ticker.PriceUSD, ticker.PriceBTC)
		}
	}
}
