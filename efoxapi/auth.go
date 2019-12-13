package efoxapi

import (
	"crypto/ecdsa"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWTToken(uid string, key *ecdsa.PrivateKey, exp time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(exp).Unix(),
	})

	return token.SignedString(key)
}
