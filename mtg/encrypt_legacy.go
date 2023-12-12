package mtg

import (
	"crypto/ed25519"

	"github.com/fox-one/mixin-sdk-go/v2/mixinnet"
)

// EncryptLegacy encrypts data with legacy aes key & iv
// Deprecated, use Encrypt instead
func EncryptLegacy(body []byte, privateKey ed25519.PrivateKey, publicKey ed25519.PublicKey) ([]byte, error) {
	key, iv, err := keyPairsToAesKeyIvLegacy(privateKey, publicKey)
	if err != nil {
		return nil, err
	}

	prefix := privateKey.Public().(ed25519.PublicKey)
	return encryptWithAesKeyIv(body, prefix, key, iv, nil)
}

// DecryptLegacy decrypts the body with the private key and public key.
// Deprecated, use Decrypt instead
func DecryptLegacy(b []byte, privateKey ed25519.PrivateKey) ([]byte, error) {
	r := NewReader(b)

	pub, err := r.Read(ed25519.PublicKeySize)
	if err != nil {
		return nil, err
	}

	data, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	key, iv, err := keyPairsToAesKeyIvLegacy(privateKey, pub)
	if err != nil {
		return nil, err
	}

	return decryptWithAseKeyIv(data, key, iv, nil)
}

// keyPairsToAesKeyIvLegacy generate aes key and iv from key pairs
// Deprecated. this function is buggy, use keyPairsToAesKeyIv instead
func keyPairsToAesKeyIvLegacy(_ ed25519.PrivateKey, publicKey ed25519.PublicKey) (key, iv []byte, err error) {
	var pri, pub mixinnet.Key
	copy(pub[:], publicKey)
	// privateKeyToCurve25519(pri, privateKey)

	point := mixinnet.KeyMultPubPriv(&pub, &pri)
	tmp := point.Bytes()
	return tmp[:16], tmp[16:], nil
}
