package telesign

import "context"

type Client struct {
	Key    string
	Secret string
}

func NewClient(key, secret string) *Client {
	return &Client{
		Key:    key,
		Secret: secret,
	}
}

func (c *Client) SendMessage(ctx context.Context, msg Message) error {
	return SendMessage(ctx, c.Key, c.Secret, msg)
}
