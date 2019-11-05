package foxapi

import "github.com/dgrijalva/jwt-go"

func ValidateToken(token string) error {
	var claim jwt.StandardClaims
	_, _ = jwt.ParseWithClaims(token, &claim, nil)
	return claim.Valid()
}
