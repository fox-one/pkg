package f1exapi

import (
	"context"
	"strconv"

	"github.com/fox-one/pkg/pagination"
	"github.com/shopspring/decimal"
)

type Order struct {
	OrderID         string          `json:"id"`
	CreatedAt       int64           `json:"created_at"`
	OrderType       string          `json:"order_type"`
	UserID          string          `json:"user_id"`
	QuoteAssetID    string          `json:"quote_asset_id"`
	BaseAssetID     string          `json:"base_asset_id"`
	Symbol          string          `json:"symbol"`
	Side            string          `json:"side"`
	Price           decimal.Decimal `json:"price"`
	RemainingAmount decimal.Decimal `json:"remaining_amount"`
	FilledAmount    decimal.Decimal `json:"filled_amount"`
	RemainingFunds  decimal.Decimal `json:"remaining_fund"`
	FilledFunds     decimal.Decimal `json:"filled_fund"`
	State           string          `json:"state"`
}

func QueryOrder(ctx context.Context, token, orderID string) (*Order, error) {
	resp, err := request(ctx).SetAuthToken(token).Get("/order/" + orderID)
	if err != nil {
		return nil, err
	}

	var order Order
	if err := decodeResponse(resp, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

type QueryOrdersInput struct {
	Symbol string
	Side   string
	Cursor string
	Limit  int
	Order  string
	State  string
	Start  int64 // 单位秒
	End    int64 // 单位秒
}

func (input *QueryOrdersInput) toParams() map[string]string {
	params := map[string]string{
		"symbol": input.Symbol,
		"cursor": input.Cursor,
	}

	if input.Side != "" {
		params["side"] = input.Side
	}

	if input.Limit > 0 {
		params["limit"] = strconv.Itoa(input.Limit)
	}

	if input.Order != "" {
		params["order"] = input.Order
	}

	if input.State != "" {
		params["state"] = input.State
	}

	if input.Start > 0 {
		params["start"] = strconv.FormatInt(input.Start, 10)
	}

	if input.End > 0 {
		params["end"] = strconv.FormatInt(input.End, 10)
	}

	return params
}

func QueryOrders(ctx context.Context, token string, input *QueryOrdersInput) ([]*Order, *pagination.Pagination, error) {
	resp, err := request(ctx).SetAuthToken(token).SetQueryParams(input.toParams()).Get("/orders")
	if err != nil {
		return nil, nil, err
	}

	var body struct {
		Pagination *pagination.Pagination `json:"pagination,omitempty"`
		Orders     []*Order               `json:"orders,omitempty"`
	}

	if err := decodeResponse(resp, &body); err != nil {
		return nil, nil, err
	}

	return body.Orders, body.Pagination, nil
}
