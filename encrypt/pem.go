package encrypt

import (
	"encoding/pem"
	"errors"

	"github.com/hashicorp/go-multierror"
)

func ParsePrivatePem(value string) (pri, pub interface{}, err error) {
	block, _ := pem.Decode([]byte(value))
	if block == nil {
		err = errors.New("decode pem failed: block is nil")
		return
	}

	if key, err2 := RSA.Parse(block.Bytes); err2 == nil {
		return key, key.Public(), nil
	} else {
		err = multierror.Append(err, err2)
	}

	if key, err2 := ECDSA.Parse(block.Bytes); err2 == nil {
		return key, key.Public(), nil
	} else {
		err = multierror.Append(err, err2)
	}

	return
}
