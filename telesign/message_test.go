package telesign

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendMessage(t *testing.T) {
	c := &Client{
		Key:    "",
		Secret: "",
	}

	err := c.SendMessage(context.Background(), Message{
		Phone:   "your phone",
		Content: "test telesign",
		Type:    "",
	})
	assert.Nil(t, err)
}
