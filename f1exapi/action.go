package f1exapi

import (
	"encoding/base64"
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/fox-one/pkg/number"
	"github.com/gofrs/uuid"
	"github.com/ugorji/go/codec"
)

const (
	OrderSideASK = "ASK"
	OrderSideBID = "BID"

	OrderTypeLimit  = "LIMIT"
	OrderTypeMarket = "MARKET"

	OrderStatePending = "PENDING"
	OrderStateDone    = "DONE"

	TransferSourceCancel = "CANCEL"
	TransferSourceRefund = "REFUND"
	TransferSourceMatch  = "MATCH"

	TradeSideMaker = "MAKER"
	TradeSideTaker = "TAKER"
)

type OrderAction struct {
	S string    // side
	A uuid.UUID // asset
	P string    // price
	T string    // type
	O uuid.UUID // order id
	M uuid.UUID // merchant
	U []byte    // user public key
}

func (action *OrderAction) toParams() map[string]interface{} {
	params := map[string]interface{}{}

	if action.S != "" {
		params["S"] = action.S
	}

	if action.A != uuid.Nil {
		params["A"] = action.A
	}

	if action.P != "" {
		params["P"] = action.P
	}

	if action.T != "" {
		params["T"] = action.T
	}

	if action.O != uuid.Nil {
		params["O"] = action.O
	}

	if action.M != uuid.Nil {
		params["M"] = action.M
	}

	if len(action.U) > 0 {
		params["U"] = action.U
	}

	return params
}

type TransferAction struct {
	S string    // source
	O uuid.UUID // order
	A uuid.UUID // asset id
	P string    // price
	C string    // category, bid or ask
}

func (action *OrderAction) Encode() (string, error) {
	memo := make([]byte, 140)
	handle := new(codec.MsgpackHandle)
	encoder := codec.NewEncoderBytes(&memo, handle)
	if err := encoder.Encode(action.toParams()); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(memo), nil
}

type PutOrderInput struct {
	Base       string `valid:"uuid,required"`
	Quote      string `valid:"uuid,required"`
	Side       string `valid:"in(ASK|BID),required"`
	Type       string `valid:"in(LIMIT|MARKET),required"`
	Price      string
	MerchantID string `valid:"uuid"`
}

type PutOrderOutput struct {
	AssetID  string
	Memo     string
	Opponent string
}

func PutOrder(input *PutOrderInput) (*PutOrderOutput, error) {
	if _, err := govalidator.ValidateStruct(input); err != nil {
		return nil, err
	}

	action := OrderAction{
		S: input.Side,
		T: input.Type,
		M: uuid.FromStringOrNil(input.MerchantID),
	}

	var out PutOrderOutput

	switch input.Side {
	case OrderSideASK:
		out.AssetID = input.Base
		action.A = uuid.FromStringOrNil(input.Quote)
	case OrderSideBID:
		out.AssetID = input.Quote
		action.A = uuid.FromStringOrNil(input.Base)
	}

	if input.Type == OrderTypeLimit {
		price := number.Decimal(input.Price)
		if !price.IsPositive() {
			return nil, errors.New("price must be positive for limit order")
		}

		action.P = price.String()
	}

	memo, err := action.Encode()
	if err != nil {
		return nil, err
	}

	out.Memo = memo
	return &out, nil
}

func ParsePutOrder(memo string) (*OrderAction, error) {
	data, err := base64.StdEncoding.DecodeString(memo)
	if err != nil {
		data, err = base64.URLEncoding.DecodeString(memo)
	}

	if err != nil {
		return nil, err
	}

	handle := new(codec.MsgpackHandle)
	decoder := codec.NewDecoderBytes(data, handle)

	var action OrderAction
	err = decoder.Decode(&action)
	return &action, err
}

func ParseTransfer(memo string) (*TransferAction, error) {
	data, err := base64.StdEncoding.DecodeString(memo)
	if err != nil {
		data, err = base64.URLEncoding.DecodeString(memo)
	}

	if err != nil {
		return nil, err
	}

	handle := new(codec.MsgpackHandle)
	decoder := codec.NewDecoderBytes(data, handle)

	var action TransferAction
	err = decoder.Decode(&action)
	return &action, err
}
