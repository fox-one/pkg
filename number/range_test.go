package number

import (
	"testing"
)

func TestChunk(t *testing.T) {
	Chunk(100, 11, func(l, r int) {
		t.Log(l, r)
	})
}
