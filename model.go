package yakassa

import (
	"context"
	"time"
)

type Currency string

const (
	CurrencyRUB Currency = "RUB"
)

type CreatePaymentRequest struct {
	Amount            Amount             `json:"amount"`
	Capture           bool               `json:"capture"`
	Confirmation      Confirmation       `json:"confirmation"`
	Description       string             `json:"description"`
	PaymentMethodData *PaymentMethodData `json:"payment_method_data,omitempty"`

	// used in op func for inject context
	Ctx            context.Context `json:"-"`
	IdempotenceKey string          `json:"-"`
}
type Amount struct {
	Value    string   `json:"value"`
	Currency Currency `json:"currency"`
}
type Confirmation struct {
	Type      string `json:"type"`
	ReturnURL string `json:"return_url"`
}
type PaymentMethodData struct {
	Type string `json:"type"`
}

type CreatePaymentResponse struct {
	ID     string `json:"id"`
	Type   string `json:"type"` // for error response
	Code   string `json:"code"` // for error response
	Status string `json:"status"`
	Paid   bool   `json:"paid"`
	Amount struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"amount"`
	Confirmation struct {
		Type            string `json:"type"`
		ConfirmationURL string `json:"confirmation_url"`
	} `json:"confirmation"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	Metadata    struct {
	} `json:"metadata"`
	Recipient struct {
		AccountID string `json:"account_id"`
		GatewayID string `json:"gateway_id"`
	} `json:"recipient"`
	Refundable bool `json:"refundable"`
	Test       bool `json:"test"`
}
