package airdropapi

import (
	"context"
	"testing"

	"github.com/fox-one/pkg/number"
	"github.com/fox-one/pkg/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPutAirdrop(t *testing.T) {
	ctx := context.Background()
	drop, err := PutAirdrop(ctx, "test", uuid.New(), uuid.New(), "test", []Target{
		{WalletID: uuid.New(), Amount: number.Decimal("10")},
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, drop.Recipient)
	t.Log(drop.Recipient)
}
