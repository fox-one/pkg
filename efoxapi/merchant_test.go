package efoxapi

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListOrderReports(t *testing.T) {
	ctx := context.Background()
	reports, next, err := ListOrderReports(ctx, "xxx", time.Now(), "", 10)
	assert.Nil(t, err)
	t.Log(reports, next)
}
