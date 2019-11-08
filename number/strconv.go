package number

import (
	"strconv"
)

func ParseInt64(s string) (int64, bool) {
	i, err := strconv.ParseInt(s, 10, 64)
	return i, err == nil
}
