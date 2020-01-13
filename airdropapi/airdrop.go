package airdropapi

import (
	"context"

	"github.com/shopspring/decimal"
)

type Target struct {
	WalletID string          `json:"wallet_id,omitempty"`
	Amount   decimal.Decimal `json:"amount,omitempty"`
}

type Airdrop struct {
	AssetID   string `json:"asset_id"`
	Amount    string `json:"amount"`
	Recipient string `json:"recipient"`
	TraceID   string `json:"trace_id"`
}

func PutAirdrop(ctx context.Context, owner, traceID, assetID, memo string, targets []Target) (*Airdrop, error) {
	resp, err := request(ctx).SetBody(map[string]interface{}{
		"owner":    owner,
		"trace_id": traceID,
		"asset_id": assetID,
		"memo":     memo,
		"targets":  targets,
	}).Post("/api/airdrop/request.json")
	if err != nil {
		return nil, err
	}

	var drop Airdrop
	if err := decodeResponse(resp, &drop); err != nil {
		return nil, err
	}

	return &drop, nil
}
