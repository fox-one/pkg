package encrypt

import (
	"crypto/rsa"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePublicPem(t *testing.T) {
	value := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDEDtIRT57TJAfmub2RsIM32jdo
8ijsds/u1fpY6hwtkC01/LFJkNTXqSwvpaO5tp86o0SlzBHdF0WxPtsKqdc8F7kQ
uHm7hUTLX0zPGRdGCsy9q/PIGlVGAFTBSVXl+grmGGZuS1CHI13L/oulBGENQOxO
8r6D1RyPjt6z0BAndQIDAQAB
-----END PUBLIC KEY-----`

	pub, err := ParsePublicPem(value)
	assert.Nil(t, err)
	assert.IsType(t, &rsa.PublicKey{}, pub)
}
