package foxapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type Asset struct {
	AssetID       string          `json:"asset_id,omitempty"`
	ChainID       string          `json:"chain_id,omitempty"`
	Name          string          `json:"name,omitempty"`
	Symbol        string          `json:"symbol,omitempty"`
	Icon          string          `json:"icon,omitempty"`
	Price         decimal.Decimal `json:"price,omitempty"`
	Change        decimal.Decimal `json:"change,omitempty"`
	PriceUSD      decimal.Decimal `json:"price_usd,omitempty"`
	ChangeUSD     decimal.Decimal `json:"change_usd,omitempty"`
	PriceBTC      decimal.Decimal `json:"price_btc,omitempty"`
	ChangeBTC     decimal.Decimal `json:"change_btc,omitempty"`
	Confirmations int             `json:"confirmations,omitempty"`

	// user asset
	Balance     decimal.Decimal `json:"balance,omitempty"`
	Destination string          `json:"destination,omitempty"`
	Tag         string          `json:"tag,omitempty"`

	Chain *Asset `json:"chain,omitempty"`
}

func ReadAssets(ctx context.Context, accessToken string) ([]*Asset, error) {
	resp, err := request(ctx).SetAuthToken(accessToken).Get("/wallet/assets")
	if err != nil {
		return nil, err
	}

	var assets []*Asset
	err = decodeResponse(resp, &assets)
	return assets, err
}

func ReadAsset(ctx context.Context, accessToken string, assetID string) (*Asset, error) {
	resp, err := request(ctx).SetAuthToken(accessToken).Get("/wallet/asset/" + assetID)
	if err != nil {
		return nil, err
	}

	var asset Asset
	err = decodeResponse(resp, &asset)
	return &asset, err
}

func SearchAssets(ctx context.Context, symbol string, fuzzy bool) ([]*Asset, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	if fuzzy {
		params["mode"] = "fuzzy"
	}

	resp, err := request(ctx).SetQueryParams(params).Get("/wallet/search-assets")
	if err != nil {
		return nil, err
	}

	var assets []*Asset
	err = decodeResponse(resp, &assets)
	return assets, err
}
