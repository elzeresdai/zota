package main

import (
	"log"
	"net/http"
	"zota_test/domain/service"
	"zota_test/infrastructure/zota"
	http2 "zota_test/interfaces/http"
	"zota_test/interfaces/http/handler"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	zotaAdapter := &zota.ZotaAdapter{}
	paymentService := service.NewPaymentService(zotaAdapter)

	depositHandler := handler.NewDepositHandler(paymentService)
	statusHandler := handler.NewStatusHandler(paymentService)

	router := http2.NewRouter(depositHandler, statusHandler)

	log.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
