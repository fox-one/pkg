package sqs

import (
	"strings"
)

func isFifoQueue(url string) bool {
	return strings.HasSuffix(url, ".fifo")
}
