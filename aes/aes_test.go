package aes

import (
	"crypto/rand"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	for _, size := range []int{16, 24, 32} {
		t.Run(fmt.Sprintf("aes key size %d", size), func(t *testing.T) {
			var (
				body = make([]byte, 255)
				key  = make([]byte, size)
				iv   = make([]byte, 16)
			)

			for range body {
				_, _ = io.ReadFull(rand.Reader, body)
				_, _ = io.ReadFull(rand.Reader, key)
				_, _ = io.ReadFull(rand.Reader, iv)

				encrypted, err := Encrypt(body, key, iv)
				require.NoError(t, err)

				decrypted, err := Decrypt(encrypted, key, iv)
				require.NoError(t, err)

				assert.Equal(t, body, decrypted)
			}
		})
	}
}

func TestDecryptEmpty(t *testing.T) {
	var (
		body = make([]byte, 0)
		key  = make([]byte, 16)
		iv   = make([]byte, 16)
	)

	_, _ = io.ReadFull(rand.Reader, body)
	_, _ = io.ReadFull(rand.Reader, key)
	_, _ = io.ReadFull(rand.Reader, iv)

	decrypted, err := Decrypt(body, key, iv)
	require.Error(t, err)
	assert.Lenf(t, decrypted, 0, "decrypted body should be empty")
}
