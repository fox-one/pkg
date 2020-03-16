package foxapi

import (
	"context"
	"strconv"

	"github.com/fox-one/pkg/pagination"
	"github.com/shopspring/decimal"
)

type Snapshot struct {
	Asset           *Asset                 `json:"asset"`
	Amount          decimal.Decimal        `json:"amount"`
	AssetID         string                 `json:"asset_id"`
	InsideMixin     bool                   `json:"inside_mixin,omitempty"`
	CreatedAt       int64                  `json:"created_at"`
	Memo            string                 `json:"memo,omitempty"`
	ExtraData       map[string]interface{} `json:"extra_data"`
	Opponent        *Opponent              `json:"opponent,omitempty"`
	OpponentID      string                 `json:"opponent_id,omitempty"`
	Receiver        string                 `json:"receiver,omitempty"`
	Sender          string                 `json:"sender,omitempty"`
	SnapshotID      string                 `json:"snapshot_id,omitempty"`
	Source          string                 `json:"source,omitempty"`
	TraceID         string                 `json:"trace_id,omitempty"`
	TransactionHash string                 `json:"transaction_hash,omitempty"`
	UserID          string                 `json:"user_id,omitempty"`
}

type Opponent struct {
	Avatar   string `json:"avatar,omitempty"`
	FullName string `json:"fullname,omitempty"`
	ID       string `json:"id,omitempty"`
}

type PageSnapshots struct {
	Pagination *pagination.Pagination `json:"pagination"`
	Snapshots  []*Snapshot            `json:"snapshots"`
}

func ReadSnapshots(context context.Context, accessToken, assetID, cursor string, limit int) (*PageSnapshots, error) {
	queryParams := make(map[string]string, 0)
	if assetID != "" {
		queryParams["asset_id"] = assetID
	}
	if cursor != "" {
		queryParams["cursor"] = cursor
	}

	if limit <= 0 {
		limit = 30
	}
	queryParams["limit"] = strconv.Itoa(limit)

	resp, err := request(context).SetAuthToken(accessToken).SetQueryParams(queryParams).Get("/wallet/snapshots")
	if err != nil {
		return nil, err
	}

	var pageSnapshots PageSnapshots

	err = decodeResponse(resp, &pageSnapshots)

	return &pageSnapshots, nil
}

func ReadSnapshot(context context.Context, accessToken, snapshotID string) (*Snapshot, error) {
	resp, err := request(context).SetAuthToken(accessToken).Get("/wallet/snapshot/" + snapshotID)
	if err != nil {
		return nil, err
	}

	var snapshot Snapshot
	err = decodeResponse(resp, &snapshot)
	return &snapshot, err
}
