package handler

import (
	"github.com/stretchr/testify/mock"
	"zota_test/domain/model"
)

type MockPaymentService struct {
	mock.Mock
}

func (m *MockPaymentService) MakeDeposit(request model.DepositRequest) (*model.DepositResponse, error) {
	args := m.Called(request)
	if args.Get(0) != nil {
		return args.Get(0).(*model.DepositResponse), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockPaymentService) CheckStatus(request model.StatusRequest) (*model.StatusResponse, error) {
	args := m.Called(request)
	if args.Get(0) != nil {
		return args.Get(0).(*model.StatusResponse), args.Error(1)
	}
	return nil, args.Error(1)
}
