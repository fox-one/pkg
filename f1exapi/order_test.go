package f1exapi

import (
	"context"
	"testing"

	"github.com/fox-one/pkg/encrypt"
	"github.com/stretchr/testify/assert"
)

const (
	merchantID  = "aa6911ca-5fe6-3061-97fe-59fa2034e420"
	merchantKey = `-----BEGIN RSA PRIVATE KEY-----
MHcCAQEEIK8E20km9juc+sza6ZKjeH5xcfTMdads4O70jjjtmJAvoAoGCCqGSM49
AwEHoUQDQgAEFB6V7yLFtyYNbo6MgC1Ljsxl6ag8EaSLKYdTRgCe7ygUuq3S1GJb
F07WMTxb/27tLZlEy09hRx842FTguLheXA==
-----END RSA PRIVATE KEY-----`
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
