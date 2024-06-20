package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"zota_test/domain/model"
	"zota_test/domain/service"
)

type DepositHandler struct {
	paymentService service.PaymentServiceInterface
}

func NewDepositHandler(paymentService service.PaymentServiceInterface) *DepositHandler {
	return &DepositHandler{paymentService: paymentService}
}

func (h *DepositHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	amount := r.URL.Query().Get("amount")
	currency := r.URL.Query().Get("currency")
	request := model.DepositRequest{
		MerchantOrderID:     generateUniqueOrderID(),
		MerchantOrderDesc:   "Test Order Description",
		OrderAmount:         amount,
		OrderCurrency:       currency,
		CustomerEmail:       "[email protected]",
		CustomerFirstName:   "John",
		CustomerLastName:    "Doe",
		CustomerAddress:     "123 Example St",
		CustomerCountryCode: "US",
		CustomerCity:        "Example City",
		CustomerZipCode:     "12345",
		CustomerPhone:       "+1234567890",
		CustomerIP:          "127.0.0.1",
		RedirectUrl:         "https://www.example-merchant.com/payment-return/",
		CallbackUrl:         "https://www.example-merchant.com/payment-callback/",
		CheckoutUrl:         "https://www.example-merchant.com/account/deposit/",
	}
	resp, err := h.paymentService.MakeDeposit(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func generateUniqueOrderID() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.FormatInt(rand.Int63(), 10)
}
