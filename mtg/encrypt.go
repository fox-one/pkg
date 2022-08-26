package mtg

import (
	"bytes"
	"crypto/ed25519"
	"crypto/md5"
	"fmt"
	"hash"

	"github.com/fox-one/pkg/encrypt"
	"golang.org/x/crypto/curve25519"
)

func Encrypt(body []byte, privateKey ed25519.PrivateKey, publicKey ed25519.PublicKey) ([]byte, error) {
	key, iv, err := keyPairsToAesKeyIv(privateKey, publicKey)
	if err != nil {
		return nil, err
	}

	prefix := privateKey.Public().(ed25519.PublicKey)
	return encryptWithAesKeyIv(body, prefix, key, iv, md5.New())
}

// Decrypt decrypts the body with the private key and public key.
func Decrypt(b []byte, privateKey ed25519.PrivateKey) ([]byte, error) {
	r := NewReader(b)

	pub, err := r.Read(ed25519.PublicKeySize)
	if err != nil {
		return nil, err
	}

	data, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	key, iv, err := keyPairsToAesKeyIv(privateKey, pub)
	if err != nil {
		return nil, err
	}

	return decryptWithAseKeyIv(data, key, iv, md5.New())
}

func keyPairsToAesKeyIv(privateKey ed25519.PrivateKey, publicKey ed25519.PublicKey) (key, iv []byte, err error) {
	var pri, pub [32]byte

	privateKeyToCurve25519(&pri, privateKey)
	if err = publicKeyToCurve25519(&pub, publicKey); err != nil {
		return
	}

	var dst []byte
	if dst, err = curve25519.X25519(pri[:], pub[:]); err != nil {
		return
	}

	if l := len(dst); l != 32 {
		err = fmt.Errorf("bad scalar multiplication length: %d, expected %d", l, 32)
		return
	}

	key, iv = dst, md5Hash(dst)
	return
}

func encryptWithAesKeyIv(body, prefix, key, iv []byte, h hash.Hash) ([]byte, error) {
	if h != nil {
		if _, err := h.Write(body); err != nil {
			return nil, err
		}

		body = append(h.Sum(nil), body...)
	}

	data, err := encrypt.Encrypt(body, key, iv)
	if err != nil {
		return nil, err
	}

	dst := make([]byte, len(prefix)+len(data))
	n := copy(dst, prefix)
	_ = copy(dst[n:], data)
	return dst, nil
}

func decryptWithAseKeyIv(data, key, iv []byte, h hash.Hash) ([]byte, error) {
	b, err := encrypt.Decrypt(data, key, iv)
	if err != nil {
		return nil, err
	}

	if h != nil {
		r := NewReader(b)
		sig, err := r.Read(h.Size())
		if err != nil {
			return nil, err
		}

		b, err = r.ReadAll()
		if err != nil {
			return nil, err
		}

		if _, err := h.Write(b); err != nil {
			return nil, err
		}

		if !bytes.Equal(h.Sum(nil), sig) {
			return nil, fmt.Errorf("invalid signature")
		}
	}

	return b, nil
}
