package service

import (
	"zota_test/domain/model"
	"zota_test/domain/repository"
)

type PaymentService struct {
	repo repository.PaymentRepositoryInterface
}

func NewPaymentService(repo repository.PaymentRepositoryInterface) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) MakeDeposit(request model.DepositRequest) (*model.DepositResponse, error) {
	return s.repo.MakeDeposit(request)
}

func (s *PaymentService) CheckStatus(request model.StatusRequest) (*model.StatusResponse, error) {
	return s.repo.CheckStatus(request)
}
