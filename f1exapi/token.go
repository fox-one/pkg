package f1exapi

import (
	"crypto/ecdsa"

	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	jwt.StandardClaims
	MerchantID string `json:"mid,omitempty"`
	UserID     string `json:"uid,omitempty"`
}

func SignToken(claim Claim, key *ecdsa.PrivateKey) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodES256, claim).SignedString(key)
}
