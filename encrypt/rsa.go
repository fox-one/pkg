package encrypt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
)

// RSA Private Key
var RSA rsaKey

type rsaKey struct{}

func (rsaKey) New() *rsa.PrivateKey {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	return key
}

func (rsaKey) Marshal(key *rsa.PrivateKey) []byte {
	return x509.MarshalPKCS1PrivateKey(key)
}

func (rsaKey) Parse(block []byte) (*rsa.PrivateKey, error) {
	return x509.ParsePKCS1PrivateKey(block)
}

// ECDSA Private Key
var ECDSA ecdsaKey

type ecdsaKey struct{}

func (ecdsaKey) New() *ecdsa.PrivateKey {
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return key
}

func (ecdsaKey) Marshal(key *ecdsa.PrivateKey) []byte {
	data, _ := x509.MarshalECPrivateKey(key)
	return data
}

func (ecdsaKey) Parse(block []byte) (*ecdsa.PrivateKey, error) {
	return x509.ParseECPrivateKey(block)
}

// Public Key
var Public publicKey

type publicKey struct{}

func (publicKey) Marshal(key interface{}) ([]byte, error) {
	return x509.MarshalPKIXPublicKey(key)
}

func (publicKey) Parse(block []byte) (interface{}, error) {
	return x509.ParsePKIXPublicKey(block)
}
