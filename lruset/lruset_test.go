package lruset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLruSet(t *testing.T) {
	t.Run("add value", func(t *testing.T) {
		set := New(2)
		set.Add("1")
		assert.True(t, set.Contains("1"), "1 should in")
	})

	t.Run("test remove value", func(t *testing.T) {
		set := New(2)
		set.Add("1")
		set.Add("2")
		set.Remove("1")
		assert.False(t, set.Contains("1"), "1 should be removed")
	})

	t.Run("discard values", func(t *testing.T) {
		set := New(2)
		set.Add("1")
		set.Add("2")
		set.Add("3")

		assert.False(t, set.Contains("1"), "1 should be discarded")
		assert.True(t, set.Contains("2"), "2 should in set")
		assert.True(t, set.Contains("3"), "3 should in set")
	})
}
