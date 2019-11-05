package foxapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserMe(t *testing.T) {
	ctx := context.Background()

	t.Run("invalid token", func(t *testing.T) {
		user, err := UserMe(ctx, "invalid token")
		assert.Nil(t, user)
		assert.True(t, IsError(err, ErrAuthFailed))
	})
}
