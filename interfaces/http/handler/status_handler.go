package handler

import (
	"encoding/json"
	"net/http"

	"zota_test/domain/model"
	"zota_test/domain/service"
)

type StatusHandler struct {
	paymentService service.PaymentServiceInterface
}

func NewStatusHandler(paymentService service.PaymentServiceInterface) *StatusHandler {
	return &StatusHandler{paymentService: paymentService}
}

func (h *StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	merchantOrderID := r.URL.Query().Get("merchantOrderID")
	request := model.StatusRequest{
		MerchantOrderID: merchantOrderID,
	}
	resp, err := h.paymentService.CheckStatus(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
