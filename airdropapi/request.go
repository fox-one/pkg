package airdropapi

import (
	"context"
	"encoding/json"

	"github.com/asaskevich/govalidator"
	"github.com/fox-one/pkg/foxerr"
	"github.com/go-resty/resty/v2"
)

const (
	Endpoint = "https://openapi.fox.one"
)

var client = resty.New().
	SetHeader("Content-Type", "application/json").
	SetHostURL(Endpoint)

func UseEndpoint(endpoint string) {
	if ok := govalidator.IsURL(endpoint); !ok {
		panic("endpoint must be valid url")
	}

	client = client.SetHostURL(endpoint)
}

func request(ctx context.Context) *resty.Request {
	return client.R().SetContext(ctx)
}

func decodeResponse(resp *resty.Response, data interface{}) error {
	var body struct {
		*foxerr.Error
	}

	if err := json.Unmarshal(resp.Body(), &body); err != nil {
		if resp.IsError() {
			return foxerr.New(resp.StatusCode(), resp.Status())
		}

		return err
	}

	if err := body.Error; err != nil && err.Code > 0 {
		return err
	}

	if data != nil {
		if err := json.Unmarshal(resp.Body(), data); err != nil {
			return err
		}
	}

	return nil
}
