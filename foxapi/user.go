package foxapi

import (
	"context"
)

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
