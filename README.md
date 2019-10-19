# Yandex kassa golang API client

Клиент для работы с АПИ яндекс кассы из go.

## Пример создания платежа

```go

shopID := "1"
secret := "secret"
yakassaClient := yakass.New(yakassa.Creds(shopID, secret))

paymentAmount = 150 // сумма платежа

paymentResponse, err := yakassaClient.CreatePayment(yakassa.PaymentRUB(paymentAmount))
if err!=nil{}

println(paymentResponse.Confirmation.ConfirmationURL) // URL по которому нужно отправить пользователя для процесса оплаты

```
