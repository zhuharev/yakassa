package yakassa_test

import (
	"testing"

	"github.com/zhuharev/yakassa"
)

func TestPaymentRUB(t *testing.T) {
	client := yakassa.New(yakassa.Creds("636263", "test_7VH6ISru_yRgd6HUxmeRgegdhbEyX6DlOLQBsu3SV0Q"), yakassa.DefaultReturnURL("http://localhost:8000"), yakassa.Verbose())
	response, err := client.CreatePayment(yakassa.PaymentRUB(100))
	if err != nil {
		t.Errorf("create payment error: %s", err)
	}
	if response == nil {
		t.Errorf("response must not be nil")
	}
	t.Logf("%+v", response)
	t.Error()
}
