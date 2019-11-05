package foxapi

import (
	"context"
)

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

func SearchWalletUser(ctx context.Context, walletID string) (*User, error) {
	resp, err := request(ctx).Get("/wallet/user/" + walletID)
	if err != nil {
		return nil, err
	}

	var user User
	if err := decodeResponse(resp, &user); err != nil {
		return nil, err
	}

	return &user, nil
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
