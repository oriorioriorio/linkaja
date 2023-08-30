package repository

import (
	"context"
	"errors"

	"github.com/marioheryanto/linkaja/go-app/model"
	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	Mock mock.Mock
}

func (r *AccountRepositoryMock) GetAccountDetail(ctx context.Context, accNumber string) (model.Account, error) {
	acc := model.Account{}

	args := r.Mock.Called(accNumber)
	if args.Get(0) == nil {
		return acc, errors.New("account tidak ditemukan")
	} else {
		return args.Get(0).(model.Account), nil
	}
}

func (r *AccountRepositoryMock) TransferBalance(ctx context.Context, params model.TransferParams) error {
	args := r.Mock.Called(params)
	if args.Get(0) == nil {
		return errors.New("saldo tidak cukup")
	} else {
		return nil
	}
}
