package atmapi

import (
	"crypto/ecdsa"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWTToken(uid string, key *ecdsa.PrivateKey, exp time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(exp).Unix(),
	})

	return token.SignedString(key)
}
