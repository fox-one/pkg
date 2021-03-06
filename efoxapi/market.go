package efoxapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type (
	Market struct {
		Depth  *Depth   `json:"depth,omitempty"`
		Trades []*Trade `json:"trades,omitempty"`
		Ticker *Ticker  `json:"ticker,omitempty"`
		Pair   *Pair    `json:"pair,omitempty"`
	}

	DepthOrder [2]decimal.Decimal
	Depth      struct {
		Version string       `json:"version,omitempty"`
		Asks    []DepthOrder `json:"asks,omitempty"`
		Bids    []DepthOrder `json:"bids,omitempty"`
	}

	Trade struct {
		ID        string          `json:"id,omitempty"`
		Timestamp int64           `json:"time,omitempty"`
		Price     decimal.Decimal `json:"price,omitempty"`
		Amount    decimal.Decimal `json:"amount,omitempty"`
		Side      string          `json:"side,omitempty"`
	}

	Ticker struct {
		Last   decimal.Decimal `json:"last,omitempty"`
		Change decimal.Decimal `json:"change,omitempty"`
		Buy    decimal.Decimal `json:"buy,omitempty"`
		Sell   decimal.Decimal `json:"sell,omitempty"`
	}

	Asset struct {
		AssetID   string                     `json:"asset_id,omitempty"`
		Logo      string                     `json:"logo,omitempty"`
		Symbol    string                     `json:"symbol,omitempty"`
		Min       decimal.Decimal            `json:"min,omitempty"`
		Minimums  map[string]decimal.Decimal `json:"minimums,omitempty"`
		Max       decimal.Decimal            `json:"max,omitempty"`
		Precision int                        `json:"precision,omitempty"`
	}

	Pair struct {
		Base           *Asset          `json:"base,omitempty"`
		Quote          *Asset          `json:"quote,omitempty"`
		CanBuy         bool            `json:"can_buy,omitempty"`
		CanSell        bool            `json:"can_sell,omitempty"`
		Symbol         string          `json:"symbol,omitempty"`
		DisplaySymbol  string          `json:"display_symbol,omitempty"`
		BasePrecision  int32           `json:"base_precision,omitempty"`
		QuotePrecision int32           `json:"quote_precision,omitempty"`
		PricePrecision int32           `json:"price_precision,omitempty"`
		State          string          `json:"state,omitempty"`
		Score          int64           `json:"score,omitempty"`
		Change         decimal.Decimal `json:"change,omitempty"`
		Price          decimal.Decimal `json:"price,omitempty"`
		Exchange       string          `json:"exchange,omitempty"`
		Strategies     []string        `json:"strategies,omitempty"`
	}
)

func ReadMarket(ctx context.Context, symbol string) (*Market, error) {
	resp, err := request(ctx).SetQueryParam("symbol", symbol).Get("/market")
	if err != nil {
		return nil, err
	}

	var market Market
	if err := decodeResponse(resp, &market); err != nil {
		return nil, err
	}

	return &market, nil
}

func ReadPairs(ctx context.Context) ([]*Pair, error) {
	resp, err := request(ctx).Get("/pairs")
	if err != nil {
		return nil, err
	}

	var pairs []*Pair
	if err := decodeResponse(resp, &pairs); err != nil {
		return nil, err
	}

	return pairs, nil
}
