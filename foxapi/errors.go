package foxapi

import (
	"errors"

	"github.com/fox-one/pkg/foxerr"
)

var (
	ErrAuthFailed = foxerr.New(1537, "user auth failed")
)

func IsError(err error, target error) bool {
	return errors.Is(err, target)
}
