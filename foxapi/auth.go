package foxapi

import (
	"context"
)

func AuthorizeToken(ctx context.Context, clientID, clientSecret, code, verifier string) (*Token, error) {
	resp, err := request(ctx).SetBody(map[string]interface{}{
		"grant_type":    "authorization_code",
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code_verifier": verifier,
		"code":          code,
	}).Post("/oauth/token")

	if err != nil {
		return nil, err
	}

	var token Token
	if err := decodeResponse(resp, &token); err != nil {
		return nil, err
	}

	return &token, nil
}
