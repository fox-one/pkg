package encrypt

import (
	"github.com/fox-one/mixin-sdk/utils"
)

const (
	aesBlockSize = 32
)

func Encrypt(data []byte, key, iv []byte) (string, error) {
	return utils.Encrypt(data, key, iv, aesBlockSize)
}

func Decrypt(v string, key, iv []byte) ([]byte, error) {
	return utils.Decrypt(v, key, iv, aesBlockSize)
}
