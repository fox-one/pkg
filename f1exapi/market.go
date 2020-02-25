package f1exapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type AssetTicker struct {
	AssetID string `json:"asset_id"`
	Symbol  string `json:"symbol"`

	Volume    decimal.Decimal `json:"volume"`
	Price     decimal.Decimal `json:"price"`
	Change    decimal.Decimal `json:"change"`
	PriceUSD  decimal.Decimal `json:"price_usd"`
	ChangeUSD decimal.Decimal `json:"change_usd"`
	PriceBTC  decimal.Decimal `json:"price_btc"`
	ChangeBTC decimal.Decimal `json:"change_btc"`
}

func ReadAssetTickers(ctx context.Context) ([]*AssetTicker, error) {
	resp, err := request(ctx).Get("/market/asset-tickers")
	if err != nil {
		return nil, err
	}

	var tickers []*AssetTicker
	if err := decodeResponse(resp, &tickers); err != nil {
		return nil, err
	}

	return tickers, nil
}
