package f1exapi

import (
	"context"
	"testing"

	"github.com/fox-one/pkg/encrypt"
	"github.com/stretchr/testify/assert"
)

const (
	merchantID  = ""
	merchantKey = ``
)

func TestQueryOrders(t *testing.T) {
	key, _, err := encrypt.ParsePrivatePem(merchantKey)
	if err != nil {
		t.Fatal(err)
	}

	token, err := SignToken(Claim{MerchantID: merchantID}, key)
	if err != nil {
		t.Fatal(err)
	}

	orders, cursor, err := QueryOrders(context.Background(), token, &QueryOrdersInput{
		Symbol: "CNBXIN",
	})

	if assert.Nil(t, err) {
		assert.Empty(t, orders)
		assert.False(t, cursor.HasNext)
	}
}
