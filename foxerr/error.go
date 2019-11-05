package foxerr

import (
	"fmt"
)

type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"msg,omitempty"`
}

func (e Error) Error() string {
	if e.Code > 0 {
		return fmt.Sprintf("%s [%d]", e.Message, e.Code)
	}

	return e.String()
}

func (e Error) String() string {
	return e.Message
}

func New(code int, msg string) error {
	e := &Error{
		Code:    code,
		Message: msg,
	}

	return e
}
