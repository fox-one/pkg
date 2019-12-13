package atmapi

import (
	"context"
	"strconv"

	"github.com/fox-one/pkg/pagination"
	"github.com/shopspring/decimal"
)

type UserTradeReport struct {
	Date             string          `json:"date"`
	UserID           string          `json:"user_id"`
	PaidAssetID      string          `json:"paid_asset_id"`
	PaidAmount       decimal.Decimal `json:"paid_amount"`
	BOXAmount        decimal.Decimal `json:"box_amount"`
	BTCAmount        decimal.Decimal `json:"btc_amount"`
	XINAmount        decimal.Decimal `json:"xin_amount"`
	EOSAmount        decimal.Decimal `json:"eos_amount"`
	USDTAmount       decimal.Decimal `json:"usdt_amount"`
	ConversionAmount decimal.Decimal `json:"conversion_amount"`
	Count            int64           `json:"count"`
	CreatedAt        int64           `json:"created_at"`
}

func ListUserTradeReport(ctx context.Context, token, date, cursor string, limit int) ([]*UserTradeReport, *pagination.Pagination, error) {
	resp, err := request(ctx).
		SetAuthToken(token).
		SetQueryParams(map[string]string{
			"date":   date,
			"cursor": cursor,
			"limit":  strconv.Itoa(limit),
		}).Get("/admin/report/user-trades")
	if err != nil {
		return nil, nil, err
	}

	var body struct {
		Reports    []*UserTradeReport     `json:"reports,omitempty"`
		Pagination *pagination.Pagination `json:"pagination,omitempty"`
	}
	if err := decodeResponse(resp, &body); err != nil {
		return nil, nil, err
	}

	return body.Reports, body.Pagination, nil
}

func ListAllUserTradeReports(ctx context.Context, token, date string) ([]*UserTradeReport, error) {
	var reports []*UserTradeReport
	var cursor string
	const limit = 100

	for {
		list, page, err := ListUserTradeReport(ctx, token, date, cursor, limit)
		if err != nil {
			return nil, err
		}

		reports = append(reports, list...)

		if !page.HasNext {
			break
		}

		cursor = page.Next()
	}

	return reports, nil
}
