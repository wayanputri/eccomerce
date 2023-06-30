package data

import (
	"belajar/bareng/app/config"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

var c coreapi.Client

func setupGlobalMidtransConfigApi() {
	midtrans.ServerKey = "SB-Mid-server-0NrVq555rAslGxfgAKi_NqMn"
	// change value to `midtrans.Production`, if you want change the env to production
	midtrans.Environment = midtrans.Sandbox
}
func GoMitransts() {
	// 1. Setup with global config
	setupGlobalMidtransConfigApi()

	// Optional: here is how if you want to set append payment notification globally
	midtrans.SetPaymentAppendNotification("https://midtrans-java.herokuapp.com/notif/append1")

	// 3. Using initialize object
	c.New("SB-Mid-server-0NrVq555rAslGxfgAKi_NqMn", midtrans.Sandbox)

	// 9. Sample request charge with credit card
	requestCreditCard()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{\n    \"masked_card\": \"451111-1117\",\n    \"bank\": \"bca\",\n    \"eci\": \"06\",\n    \"channel_response_code\": \"7\",\n    \"channel_response_message\": \"Denied\",\n    \"transaction_time\": \"2021-06-08 15:49:54\",\n    \"gross_amount\": \"100000.00\",\n    \"currency\": \"IDR\",\n      \"payment_type\": \"credit_card\",\n    \"signature_key\": \"76fe68ed1b7040c7c329356c1cd47819be3ccb8b056376ff3488bfa9af1db52a85ded0501b2dab1de56e5852982133a9ef7a47c54222abbe72288c2c4f591a71\",\n    \"status_code\": \"202\",\n    \"transaction_id\": \"36f3687e-05d4-4879-a428-fd6d1ffb786e\",\n    \"transaction_status\": \"deny\",\n    \"fraud_status\": \"challenge\",\n    \"status_message\": \"Success, transaction is found\",\n    \"merchant_id\": \"G812785002\",\n    \"card_type\": \"credit\"\n}"))
	notification(w, r)
}

func requestCreditCard() {
	var c = coreapi.Client{}
	cfg := config.InitConfig()
	midtrans.ServerKey = cfg.KEY_SERVER_MIDTRANS
	authString := encodeAuthString(midtrans.ServerKey, "")
	fmt.Println("AUTH_STRING:", authString)
	midtrans.Environment = midtrans.Sandbox

	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeCreditCard,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "12345",
			GrossAmt: 200000,
		},
		CreditCard: &coreapi.BankTransferDetails{

		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "ITEM1",
				Price: 200000,
				Qty:   1,
				Name:  "Someitem",
			},
		},
	}

	res, _ := c.ChargeTransaction(chargeReq)
	fmt.Println(res)
}

func encodeAuthString(username, password string) string {
	auth := username + ":" + password
	authBytes := []byte(auth)
	encodedAuth := base64.StdEncoding.EncodeToString(authBytes)
	return encodedAuth
}

// notification : Midtrans-Go simple sample HTTP Notification handling
func notification(w http.ResponseWriter, r *http.Request) {
	// 1. Initialize empty map
	var notificationPayload map[string]interface{}

	// 2. Parse JSON request body and use it to set json to payload
	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	if err != nil {
		// do something on error when decode
		return
	}
	// 3. Get order-id from payload
	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		// do something when key `order_id` not found
		return
	}

	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, e := c.CheckTransaction(orderId)
	if e != nil {
		http.Error(w, e.GetMessage(), http.StatusInternalServerError)
		return
	} else {
		if transactionStatusResp != nil {
			// 5. Do set transaction status based on response from check transaction status
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					// TODO set transaction status on your database to 'success'
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				// TODO set transaction status on your databaase to 'success'
			} else if transactionStatusResp.TransactionStatus == "deny" {
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				// TODO set transaction status on your databaase to 'failure'
			} else if transactionStatusResp.TransactionStatus == "pending" {
				// TODO set transaction status on your databaase to 'pending' / waiting payment
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("ok"))
}