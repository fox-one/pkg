package f1exapi

import (
	"testing"

	"github.com/fox-one/pkg/encrypt"
	"github.com/stretchr/testify/assert"
)

func TestEncodeAuthMemo(t *testing.T) {
	key := encrypt.ECDSA.New()
	memo, err := EncodeAuthMemo(key.Public())
	assert.Nil(t, err)
	t.Log(memo)
	assert.True(t, len(memo) <= 140)
}
