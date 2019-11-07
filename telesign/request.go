package telesign

import (
	"context"

	"github.com/go-resty/resty/v2"
)

var httpClient *resty.Client

func request(ctx context.Context) *resty.Request {
	if httpClient == nil {
		httpClient = resty.New().SetHostURL("https://rest-ww.telesign.com/v1")
	}

	return httpClient.R().SetContext(ctx)
}
