package f1exapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type DepthOrder []decimal.Decimal

func (order DepthOrder) Values() (price, amount decimal.Decimal) {
	for idx, value := range order {
		switch idx {
		case 0:
			price = value
		case 1:
			amount = value
		default:
			return
		}
	}

	return
}

type Depth struct {
	Asks []DepthOrder `json:"asks,omitempty"`
	Bids []DepthOrder `json:"bids,omitempty"`
	// version
	LastUpdateID int64 `json:"last_update_id,omitempty"`
}

func ReadDepth(ctx context.Context, symbol string) (*Depth, error) {
	r, err := request(ctx).SetQueryParam("symbol", symbol).Get("/depth")
	if err != nil {
		return nil, err
	}

	var depth Depth
	if err := decodeResponse(r, &depth); err != nil {
		return nil, err
	}

	return &depth, nil
}
