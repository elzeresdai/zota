package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"zota_test/domain/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDepositHandler(t *testing.T) {
	mockService := new(MockPaymentService)
	depositHandler := NewDepositHandler(mockService)

	depositResponse := &model.DepositResponse{Code: "200"}

	mockService.On("MakeDeposit", mock.MatchedBy(func(req model.DepositRequest) bool {
		return req.OrderAmount == "100" && req.OrderCurrency == "USD"
	})).Return(depositResponse, nil)

	req := httptest.NewRequest("POST", "/deposit?amount=100&currency=USD", nil)
	w := httptest.NewRecorder()

	depositHandler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp model.DepositResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	assert.Equal(t, "200", resp.Code)
	mockService.AssertExpectations(t)
}
