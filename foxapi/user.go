package foxapi

import (
	"context"
)

type User struct {
	ID          string `json:"id,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	Name        string `json:"fullname,omitempty"`
	Language    string `json:"language,omitempty"`
	PhoneCode   string `json:"phone_code,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

func UserMe(ctx context.Context, accessToken string) (*User, error) {
	resp, err := request(ctx).SetAuthToken(accessToken).Get("/account/me")
	if err != nil {
		return nil, err
	}

	var user User
	if err := decodeResponse(resp, &user); err != nil {
		return nil, err
	}

	return &user, nil
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
