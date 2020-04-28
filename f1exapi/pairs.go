package f1exapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type Pair struct {
	Symbol          string          `json:"symbol,omitempty"`
	Status          string          `json:"status,omitempty"`
	Logo            string          `json:"logo,omitempty"`
	BaseAsset       string          `json:"base_asset,omitempty"`
	BaseAssetID     string          `json:"base_asset_id,omitempty"`
	QuoteAsset      string          `json:"quote_asset,omitempty"`
	QuoteAssetID    string          `json:"quote_asset_id,omitempty"`
	AmountPrecision int32           `json:"amount_precision,omitempty"`
	FundPrecision   int32           `json:"fund_precision,omitempty"`
	PricePrecision  int32           `json:"price_precision,omitempty"`
	BaseMinAmount   decimal.Decimal `json:"base_min_amount,omitempty"`
	BaseMaxAmount   decimal.Decimal `json:"base_max_amount,omitempty"`
	QuoteMinAmount  decimal.Decimal `json:"quote_min_amount,omitempty"`
	QuoteMaxAmount  decimal.Decimal `json:"quote_max_amount,omitempty"`
}

func ReadPairs(ctx context.Context) ([]*Pair, error) {
	r, err := request(ctx).Get("/market/pairs")
	if err != nil {
		return nil, err
	}

	var pairs []*Pair
	if err := decodeResponse(r, &pairs); err != nil {
		return nil, err
	}

	return pairs, nil
}
