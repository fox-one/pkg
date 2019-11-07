package foxapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchWalletUser(t *testing.T) {
	ctx := context.Background()

	t.Run("search mixin user", func(t *testing.T) {
		const id = "8017d200-7870-4b82-b53f-74bae1d2dad7"
		user, err := SearchWalletUser(ctx, id)
		assert.Nil(t, err)
		assert.Equal(t, "yiplee", user.Name)
		assert.Equal(t, id, user.ID)
	})

	t.Run("search fox user", func(t *testing.T) {
		const id = "1dbf7355-ecad-322d-899a-57c3ba72fcf2"
		user, err := SearchWalletUser(ctx, id)
		assert.Nil(t, err)
		assert.Equal(t, "yiplee", user.Name)
		assert.Equal(t, id, user.ID)
	})
}
