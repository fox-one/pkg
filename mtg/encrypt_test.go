package mtg

import (
	"crypto/ed25519"
	"crypto/rand"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	clientPublic, clientPrivate, _ := ed25519.GenerateKey(rand.Reader)
	serverPublic, serverPrivate, _ := ed25519.GenerateKey(rand.Reader)

	zeroNil := make([]byte, 16)

	t.Run("keyPairsToAesKeyIv", func(t *testing.T) {
		t.Run("current version", func(t *testing.T) {
			key1, iv1, err := keyPairsToAesKeyIv(clientPrivate, serverPublic)
			require.NoError(t, err)

			key2, iv2, err := keyPairsToAesKeyIv(serverPrivate, clientPublic)
			require.NoError(t, err)

			assert.Equal(t, key1, key2)
			assert.Equal(t, iv1, iv2)
			assert.Lenf(t, iv1, 16, "aes iv should be 16 bytes")
			assert.NotEqual(t, zeroNil, iv1, "iv should not be zero nil")
		})

		t.Run("legacy version", func(t *testing.T) {
			key1, iv1, err := keyPairsToAesKeyIvLegacy(clientPrivate, serverPublic)
			require.NoError(t, err)

			key2, iv2, err := keyPairsToAesKeyIvLegacy(serverPrivate, clientPublic)
			require.NoError(t, err)

			assert.Equal(t, key1, key2)
			assert.Equal(t, iv1, iv2)

			assert.Lenf(t, iv1, 16, "aes iv should be 16 bytes")
			assert.Equal(t, zeroNil, iv1, "legacy iv should be zero nil")
		})
	})

	body := make([]byte, 100)
	_, _ = io.ReadFull(rand.Reader, body)

	encryptedData, err := Encrypt(body, clientPrivate, serverPublic)
	require.NoError(t, err)

	decryptedData, err := Decrypt(encryptedData, serverPrivate)
	require.NoError(t, err)

	assert.Equal(t, body, decryptedData)

	t.Run("decrypt legacy data", func(t *testing.T) {
		legacyEncryptedData, err := EncryptLegacy(body, clientPrivate, serverPublic)
		require.NoError(t, err)

		assert.NotEqual(t, encryptedData, legacyEncryptedData)
		assert.Equal(t, encryptedData[:32], legacyEncryptedData[:32], "should have same prefix")

		decryptedData, err := DecryptLegacy(legacyEncryptedData, serverPrivate)
		require.NoError(t, err)

		assert.Equal(t, body, decryptedData)
	})
}
