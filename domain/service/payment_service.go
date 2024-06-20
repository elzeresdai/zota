package service

import (
	"zota_test/domain/model"
	"zota_test/domain/repository"
)

type PaymentServiceInterface interface {
	MakeDeposit(request model.DepositRequest) (*model.DepositResponse, error)
	CheckStatus(request model.StatusRequest) (*model.StatusResponse, error)
}

type paymentService struct {
	repo repository.PaymentRepositoryInterface
}

func NewPaymentService(repo repository.PaymentRepositoryInterface) PaymentServiceInterface {
	return &paymentService{repo: repo}
}

func (s *paymentService) MakeDeposit(request model.DepositRequest) (*model.DepositResponse, error) {
	return s.repo.MakeDeposit(request)
}

func (s *paymentService) CheckStatus(request model.StatusRequest) (*model.StatusResponse, error) {
	return s.repo.CheckStatus(request)
}
