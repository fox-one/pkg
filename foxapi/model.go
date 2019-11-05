package foxapi

import (
	"github.com/shopspring/decimal"
)

type User struct {
	ID          string `json:"id,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	Name        string `json:"fullname,omitempty"`
	Language    string `json:"language,omitempty"`
	PhoneCode   string `json:"phone_code,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type Token struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
	Scope        string `json:"scope,omitempty"`
	TokenType    string `json:"token_type,omitempty"` // bearer
}

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
