package efoxapi

import (
	"context"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

type (
	OrderReport struct {
		ID             string          `json:"id,omitempty"`
		CreatedAt      int64           `json:"created_at,omitempty"`
		Date           string          `json:"date,omitempty"`
		UserID         string          `json:"user_id,omitempty"`
		MerchantID     string          `json:"merchant_id,omitempty"`
		BrokerID       string          `json:"broker_id,omitempty"`
		Symbol         string          `json:"symbol,omitempty"` // BTCUSDT
		Side           string          `json:"side,omitempty"`   // BID or ASK
		FilledAmount   decimal.Decimal `json:"filled_amount,omitempty"`
		ObtainedAmount decimal.Decimal `json:"obtained_amount,omitempty"`
		FeeAmount      decimal.Decimal `json:"fee_amount,omitempty"`
		FeeAsset       string          `json:"fee_asset,omitempty"`
		Count          int64           `json:"count,omitempty"`
	}
)

func ListOrderReports(ctx context.Context, token string, date time.Time, cursor string, limit int) ([]*OrderReport, string, error) {
	resp, err := request(ctx).
		SetAuthToken(token).
		SetQueryParams(map[string]string{
			"date":   date.Format("2006-01-02"),
			"cursor": cursor,
			"limit":  strconv.Itoa(limit),
		}).Get("/order-reports")
	if err != nil {
		return nil, "", err
	}

	var body struct {
		Reports    []*OrderReport `json:"reports,omitempty"`
		Pagination struct {
			HasNext    bool   `json:"has_next,omitempty"`
			NextCursor string `json:"next_cursor,omitempty"`
		} `json:"pagination,omitempty"`
	}
	if err := decodeResponse(resp, &body); err != nil {
		return nil, "", err
	}

	next := ""
	if body.Pagination.HasNext {
		next = body.Pagination.NextCursor
	}

	return body.Reports, next, nil
}

func ListAllOrderReports(ctx context.Context, token string, date time.Time) ([]*OrderReport, error) {
	var (
		reports []*OrderReport
		cursor  string
		limit   = 100
	)

	for {
		list, next, err := ListOrderReports(ctx, token, date, cursor, limit)
		if err != nil {
			return nil, err
		}

		reports = append(reports, list...)

		if next == "" {
			break
		}

		cursor = next
	}

	return reports, nil
}
