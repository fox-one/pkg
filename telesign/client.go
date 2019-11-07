package telesign

import (
	"context"

	"github.com/go-resty/resty/v2"
)

var client *resty.Client

func request(ctx context.Context) *resty.Request {
	if client == nil {
		client = resty.New().SetHostURL("https://rest-ww.telesign.com/v1")
	}

	return client.R().SetContext(ctx)
}

type Client struct {
	Key    string
	Secret string
}
