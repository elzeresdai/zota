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

func TestStatusHandler(t *testing.T) {
	mockService := new(MockPaymentService)
	statusHandler := NewStatusHandler(mockService)

	statusResponse := &model.StatusResponse{Code: "200"}

	mockService.On("CheckStatus", mock.MatchedBy(func(req model.StatusRequest) bool {
		return req.MerchantOrderID == "test-order-id"
	})).Return(statusResponse, nil)

	req := httptest.NewRequest("GET", "/status?merchantOrderID=test-order-id", nil)
	w := httptest.NewRecorder()

	statusHandler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp model.StatusResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	assert.Equal(t, "200", resp.Code)
	mockService.AssertExpectations(t)
}
