package mtg

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
)

func DecodePrivateKey(s string) (ed25519.PrivateKey, error) {
	b := DecodeBase64(s)
	if len(b) != ed25519.PrivateKeySize {
		return nil, fmt.Errorf("invalid private key")
	}

	return b, nil
}

func DecodePublicKey(s string) (ed25519.PublicKey, error) {
	b := DecodeBase64(s)
	if len(b) != ed25519.PublicKeySize {
		return nil, fmt.Errorf("invalid public key")
	}

	return b, nil
}

func DecodeBase64(s string) []byte {
	if b, err := base64.StdEncoding.DecodeString(s); err == nil {
		return b
	}

	if b, err := base64.URLEncoding.DecodeString(s); err == nil {
		return b
	}

	return []byte(s)
}
