package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseValues(t *testing.T) {
	const input = "BTC=0.1&BTC=0.1&EOS=0.2"
	v, err := ParseValues(input)
	assert.Nil(t, err)
	assert.Equal(t, "0.2", v.Get("BTC").String())
	assert.Equal(t, "0.2", v.Get("EOS").String())
	assert.Equal(t, "BTC=0.2&EOS=0.2", v.Encode())
}
