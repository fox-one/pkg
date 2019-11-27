package f1exapi

import (
	"crypto"
	"encoding/base64"

	"github.com/fox-one/pkg/encrypt"
	"github.com/ugorji/go/codec"
)

func EncodeAuthMemo(key crypto.PublicKey) (string, error) {
	bytes, err := encrypt.Public.Marshal(key)
	if err != nil {
		return "", err
	}

	data := map[string]interface{}{
		"U": bytes,
	}

	memo := make([]byte, 140)
	handle := new(codec.MsgpackHandle)
	encoder := codec.NewEncoderBytes(&memo, handle)
	if err := encoder.Encode(data); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(memo), nil
}
