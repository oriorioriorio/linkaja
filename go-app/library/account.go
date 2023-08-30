package library

import (
	"context"
	"net/http"
	"time"

	"github.com/marioheryanto/linkaja/go-app/helper"
	"github.com/marioheryanto/linkaja/go-app/model"
	"github.com/marioheryanto/linkaja/go-app/repository"
)

type AccountLibrary struct {
	repo      repository.AccountRepositoryInterface
	validator *helper.Validator
}

type AccountLibraryInterface interface {
	CheckBalance(ctx context.Context, accNumber string) (model.Account, error)
	Transfer(ctx context.Context, params model.TransferParams) error
}

func NewAccountLibrary(repo repository.AccountRepositoryInterface, validator *helper.Validator) AccountLibraryInterface {
	return &AccountLibrary{
		repo:      repo,
		validator: validator,
	}
}

func (l *AccountLibrary) CheckBalance(ctx context.Context, accNumber string) (model.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	acc, err := l.repo.GetAccountDetail(ctx, accNumber)
	if err != nil {
		return acc, err
	}

	return acc, nil
}

func (l *AccountLibrary) Transfer(ctx context.Context, params model.TransferParams) error {
	err := l.validator.ValidateStruct(params)
	if err != nil {
		return err
	}

	fromAcc, err := l.repo.GetAccountDetail(ctx, params.FromAccountNumber)
	if err != nil {
		return err
	}

	if fromAcc.Balance < params.Amount {
		return helper.NewServiceError(http.StatusBadRequest, "saldo tidak cukup")
	}

	_, err = l.repo.GetAccountDetail(ctx, params.ToAccountNumber)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	err = l.repo.TransferBalance(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
