package foxapi

import (
	"context"
)

type Token struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	Scope        string `json:"scope,omitempty"`
	TokenType    string `json:"token_type,omitempty"` // bearer
}

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

// RefreshToken return new access token by refresh token
// old access token will still be alive in transation minutes
func RefreshToken(ctx context.Context, refreshToken string, transation int) (*Token, error) {
	resp, err := request(ctx).SetBody(map[string]interface{}{
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
		"transation":    transation,
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
