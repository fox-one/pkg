package efoxapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadMarket(t *testing.T) {
	market, err := ReadMarket(context.Background(), "BTCUSDT")
	if assert.Nil(t, err) {
		assert.Equal(t, "BTCUSDT", market.Pair.Symbol)
		assert.NotEmpty(t, market.Trades)
		assert.True(t, market.Ticker.Last.IsPositive())
	}
}
