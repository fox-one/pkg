package f1exapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadDepth(t *testing.T) {
	d, err := ReadDepth(context.Background(), "BOXUSDT")
	if assert.Nil(t, err) {
		for _, order := range d.Asks {
			price, amount := order.Values()
			assert.True(t, price.IsPositive())
			assert.True(t, amount.IsPositive())
		}

		for _, order := range d.Bids {
			price, amount := order.Values()
			assert.True(t, price.IsPositive())
			assert.True(t, amount.IsPositive())
		}
	}
}
