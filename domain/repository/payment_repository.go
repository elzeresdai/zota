package repository

import "zota_test/domain/model"

type PaymentRepositoryInterface interface {
	MakeDeposit(request model.DepositRequest) (*model.DepositResponse, error)
	CheckStatus(request model.StatusRequest) (*model.StatusResponse, error)
}
