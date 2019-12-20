package scannerapi

import (
	"context"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

type UserBalance struct {
	UserID  string          `json:"user_id"`
	Balance decimal.Decimal `json:"balance"`
}

func ScanAssetBalances(ctx context.Context, assetID string, timestampInNano int64) ([]*UserBalance, error) {
	if timestampInNano == 0 {
		timestampInNano = time.Now().UnixNano()
	}

	resp, err := request(ctx).
		SetQueryParams(map[string]string{
			"asset_id":  assetID,
			"timestamp": strconv.Itoa(int(timestampInNano)),
		}).Get("/scan-assets")

	if err != nil {
		return nil, err
	}

	var balances []*UserBalance
	if err := decodeResponse(resp, &balances); err != nil {
		return nil, err
	}

	return balances, nil
}
