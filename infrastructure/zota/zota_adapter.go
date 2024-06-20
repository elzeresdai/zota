package zota

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"os"
	"zota_test/domain/model"

	"strings"
)

type ZotaAdapter struct{}

func (za *ZotaAdapter) MakeDeposit(request model.DepositRequest) (*model.DepositResponse, error) {
	url := os.Getenv("BASE_URL") + "/api/v1/deposit/request/" + os.Getenv("ENDPOINT_ID")

	signature := generateSignature(request)
	request.Signature = signature

	jsonReqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Secret-Key", os.Getenv("API_SECRET_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var depositResp model.DepositResponse
	if err := json.NewDecoder(resp.Body).Decode(&depositResp); err != nil {
		return nil, err
	}

	return &depositResp, nil
}

func (za *ZotaAdapter) CheckStatus(request model.StatusRequest) (*model.StatusResponse, error) {
	url := os.Getenv("BASE_URL") + "/api/v1/deposit/status/" + os.Getenv("ENDPOINT_ID")

	request.Signature = generateSignatureForStatus(request)

	jsonReqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Secret-Key", os.Getenv("API_SECRET_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var statusResp model.StatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&statusResp); err != nil {
		return nil, err
	}

	return &statusResp, nil
}

func generateSignature(request model.DepositRequest) string {
	rawData := []string{
		request.MerchantOrderID,
		request.OrderAmount,
		request.OrderCurrency,
		os.Getenv("MERCHANT_ID"),
		os.Getenv("API_SECRET_KEY"),
	}
	hash := sha256.Sum256([]byte(strings.Join(rawData, ":")))
	return hex.EncodeToString(hash[:])
}

func generateSignatureForStatus(request model.StatusRequest) string {
	rawData := []string{
		request.MerchantOrderID,
		os.Getenv("MERCHANT_ID"),
		os.Getenv("API_SECRET_KEY"),
	}
	hash := sha256.Sum256([]byte(strings.Join(rawData, ":")))
	return hex.EncodeToString(hash[:])
}
