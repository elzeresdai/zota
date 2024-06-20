package model

type DepositRequest struct {
	MerchantOrderID     string `json:"merchantOrderID"`
	MerchantOrderDesc   string `json:"merchantOrderDesc"`
	OrderAmount         string `json:"orderAmount"`
	OrderCurrency       string `json:"orderCurrency"`
	CustomerEmail       string `json:"customerEmail"`
	CustomerFirstName   string `json:"customerFirstName"`
	CustomerLastName    string `json:"customerLastName"`
	CustomerAddress     string `json:"customerAddress"`
	CustomerCountryCode string `json:"customerCountryCode"`
	CustomerCity        string `json:"customerCity"`
	CustomerZipCode     string `json:"customerZipCode"`
	CustomerPhone       string `json:"customerPhone"`
	CustomerIP          string `json:"customerIP"`
	RedirectUrl         string `json:"redirectUrl"`
	CallbackUrl         string `json:"callbackUrl"`
	CheckoutUrl         string `json:"checkoutUrl"`
	Signature           string `json:"signature"`
}

type DepositResponse struct {
	Code string `json:"code"`
	Data struct {
		DepositUrl      string `json:"depositUrl"`
		MerchantOrderID string `json:"merchantOrderID"`
		OrderID         string `json:"orderID"`
	} `json:"data"`
	Message string `json:"message"`
}

type StatusRequest struct {
	MerchantOrderID string `json:"merchantOrderID"`
	Signature       string `json:"signature"`
}

type StatusResponse struct {
	Code string `json:"code"`
	Data struct {
		Status          string `json:"status"`
		MerchantOrderID string `json:"merchantOrderID"`
		OrderID         string `json:"orderID"`
	} `json:"data"`
	Message string `json:"message"`
}
