package yakassa

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const (
	defaultHTTPClientTimeout = time.Minute

	defaultAPIBaseURL = "https://payment.yandex.net/api/v3/"
)

type Logger interface {
	Printf(format string, args ...interface{})
}

type simpleLogger struct {
	log.Logger
}

type Client struct {
	httpClient *http.Client
	logger     Logger

	baseURL   string
	shopID    string
	secret    string
	returnURL string
}

// New return new client
func New(ops ...ClientOp) *Client {
	c := &Client{
		httpClient: &http.Client{
			Timeout: defaultHTTPClientTimeout,
		},
		baseURL: defaultAPIBaseURL,
	}
	for _, op := range ops {
		op(c)
	}
	return c
}

// CreatePayment create new payment
func (c *Client) CreatePayment(ops ...CreatePaymentOp) (*CreatePaymentResponse, error) {
	paymentRequest := CreatePaymentRequest{
		Capture:        true,
		IdempotenceKey: uuid.New().String(),
		Confirmation: Confirmation{
			Type:      "redirect",
			ReturnURL: c.returnURL,
		},
	}
	for _, op := range ops {
		op(&paymentRequest)
	}

	bts, err := json.Marshal(paymentRequest)
	if err != nil {
		return nil, err
	}

	if paymentRequest.Ctx == nil {
		paymentRequest.Ctx = context.Background()
	}

	req, err := http.NewRequestWithContext(paymentRequest.Ctx, "POST", c.baseURL+"payments", bytes.NewReader(bts))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.shopID, c.secret)
	req.Header.Set("Idempotence-Key", paymentRequest.IdempotenceKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	rdr := io.TeeReader(resp.Body, &buf)

	var paymentResponse CreatePaymentResponse
	err = json.NewDecoder(rdr).Decode(&paymentResponse)
	if c.logger != nil {
		c.logger.Printf("response body: %s", buf.Bytes())
	}
	if err != nil {
		return nil, fmt.Errorf("decode response err: %w body: %s", err, buf.String())
	}

	if paymentResponse.Type == "error" {
		return nil, fmt.Errorf("%s", paymentResponse.Code)
	}

	return &paymentResponse, nil
}
