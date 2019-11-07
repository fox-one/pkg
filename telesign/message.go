package telesign

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/asaskevich/govalidator"
)

const (
	MessageTypeARN = "ARN"
	MessageTypeOTP = "OTP" // one-time password
)

type Message struct {
	Phone   string `json:"phone,omitempty" valid:"required"`
	Content string `json:"content,omitempty" valid:"required"`
	Type    string `json:"type,omitempty" valid:"in(ARN|OTP)"` // default is OTP
}

func SendMessage(ctx context.Context, key, secret string, msg Message) error {
	if msg.Type == "" {
		msg.Type = MessageTypeOTP
	}

	if _, err := govalidator.ValidateStruct(msg); err != nil {
		return err
	}

	resp, err := request(ctx).
		SetBasicAuth(key, secret).
		SetFormData(map[string]string{
			"phone_number": msg.Phone,
			"message":      msg.Content,
			"message_type": msg.Type,
		}).Post("/messaging")
	if err != nil {
		return err
	}

	if resp.IsSuccess() {
		return nil
	}

	code, description := resp.StatusCode(), resp.Status()

	var body struct {
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
	}

	if err := json.Unmarshal(resp.Body(), &body); err == nil && body.Status.Code > 0 {
		code, description = body.Status.Code, body.Status.Description
	}

	return fmt.Errorf("telesign: %d %s", code, description)
}
