package yakassa

import (
	"context"
	"log"
	"os"
	"strconv"
)

type ClientOp func(*Client)

func Creds(shopID, secret string) ClientOp {
	return func(c *Client) {
		c.shopID = shopID
		c.secret = secret
	}
}

func Verbose() ClientOp {
	return func(c *Client) {
		c.logger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
	}
}

func ShopID(id string) ClientOp {
	return func(c *Client) {
		c.shopID = id
	}
}

func Secret(secret string) ClientOp {
	return func(c *Client) {
		c.secret = secret
	}
}

func DefaultReturnURL(returnURL string) ClientOp {
	return func(c *Client) {
		c.returnURL = returnURL
	}
}

type CreatePaymentOp func(*CreatePaymentRequest)

func CreatePaymentCtx(ctx context.Context) CreatePaymentOp {
	return func(req *CreatePaymentRequest) {
		req.Ctx = ctx
	}
}

func PaymentRUB(amount int) CreatePaymentOp {
	return func(req *CreatePaymentRequest) {
		req.Amount = Amount{
			Currency: CurrencyRUB,
			Value:    strconv.Itoa(amount),
		}
	}
}

func PaymentMetadata(data map[string]string) CreatePaymentOp {
	return func(req *CreatePaymentRequest) {
		req.Metadata = data
	}
}
