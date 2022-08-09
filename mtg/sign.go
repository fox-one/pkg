package mtg

import (
	"crypto/ed25519"
)

func Sign(body []byte, privateKey ed25519.PrivateKey) []byte {
	return ed25519.Sign(privateKey, body)
}

func Verify(body, sig []byte, publicKey ed25519.PublicKey) bool {
	return ed25519.Verify(publicKey, body, sig)
}

func Pack(body, sig []byte) []byte {
	b := make([]byte, len(body)+len(sig))
	n := copy(b, sig)
	copy(b[n:], body)
	return b
}

func Unpack(b []byte) (body, sig []byte, err error) {
	r := NewReader(b)
	sig, err = r.Read(ed25519.SignatureSize)
	if err != nil {
		return
	}

	body, err = r.ReadAll()
	return
}
