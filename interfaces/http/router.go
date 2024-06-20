package http

import (
	"net/http"

	"zota_test/interfaces/http/handler"
)

func NewRouter(depositHandler *handler.DepositHandler, statusHandler *handler.StatusHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/deposit", depositHandler)
	mux.Handle("/status", statusHandler)

	return mux
}
