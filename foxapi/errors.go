package foxapi

import (
	"github.com/fox-one/pkg/foxerr"
)

var (
	registerErrors = map[int]error{}

	ErrAuthFailed = newFoxErr(1537, "user auth failed")
)

func newFoxErr(code int, msg string) error {
	err := foxerr.New(code, msg)
	registerErrors[code] = err
	return err
}

func convertError(err *foxerr.Error) error {
	if v, ok := registerErrors[err.Code]; ok {
		return v
	}

	return err
}
