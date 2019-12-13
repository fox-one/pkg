package atmapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListUserTradeReport(t *testing.T) {
	ctx := context.Background()
	reports, next, err := ListUserTradeReport(ctx, "xxx", "2019-12-12", "", 100)
	assert.Nil(t, err)
	t.Log(reports, next)
}
