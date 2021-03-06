package foxerr

import (
	"errors"
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

func (e *Error) Is(target error) bool {
	if e == nil {
		return target == nil
	}

	return Match(target, e.Code)
}

func Match(err error, code int) bool {
	var target *Error
	if errors.As(err, &target) {
		return target.Code == code
	}

	return false
}
