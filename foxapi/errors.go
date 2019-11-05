package foxapi

import (
	"github.com/fox-one/pkg/foxerr"
)

var (
	ErrAuthFailed = foxerr.New(1537, "user auth failed")
)
