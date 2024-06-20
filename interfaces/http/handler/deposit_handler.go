package handler

import (
	"encoding/json"
	"net/http"

	"zota_test/domain/model"
	"zota_test/domain/service"
)

type DepositHandler struct {
	paymentService *service.PaymentService
}

func NewDepositHandler(paymentService *service.PaymentService) *DepositHandler {
	return &DepositHandler{paymentService: paymentService}
}

func (h *DepositHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	amount := r.URL.Query().Get("amount")
	currency := r.URL.Query().Get("currency")
	request := model.DepositRequest{
		MerchantOrderID:     "unique-order-id",
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
	json.NewEncoder(w).Encode(resp)
}
