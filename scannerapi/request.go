package scannerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/fox-one/pkg/foxerr"
	"github.com/go-resty/resty/v2"
	"github.com/gofrs/uuid"
)

const (
	Endpoint = "https://api.fox.one"
)

var client = resty.New().
	SetHeader("Content-Type", "application/json").
	SetHostURL(Endpoint).
	SetTimeout(10 * time.Second).
	SetPreRequestHook(func(c *resty.Client, r *http.Request) error {

		ctx := r.Context()
		if app, ok := ctx.Value(contextKeyApp).(*App); ok {
			params := r.URL.Query()
			params.Add("app_id", app.AppID)
			params.Add("nc", uuid.Must(uuid.NewV4()).String())
			params.Add("ts", strconv.Itoa(int(time.Now().Unix())))
			r.URL.RawQuery = params.Encode()
			if err := app.Auth(r); err != nil {
				return err
			}
		}

		return nil
	})

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
		Data json.RawMessage `json:"data,omitempty"`
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
		if err := json.Unmarshal(body.Data, data); err != nil {
			return err
		}
	}

	return nil
}
