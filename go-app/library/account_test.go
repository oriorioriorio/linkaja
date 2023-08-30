package library

import (
	"context"
	"testing"

	"github.com/marioheryanto/linkaja/go-app/helper"
	"github.com/marioheryanto/linkaja/go-app/model"
	"github.com/marioheryanto/linkaja/go-app/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCheckBalance(t *testing.T) {
	repo := &repository.AccountRepositoryMock{Mock: mock.Mock{}}
	lib := NewAccountLibrary(repo, helper.NewValidator())

	testCase := []struct {
		Name string
		Exp  model.Account
	}{
		{
			Name: "found",
			Exp:  model.Account{AccountNumber: "1", CustomerName: "Amir", Balance: 1},
		},
		{
			Name: "not found",
			Exp:  model.Account{},
		},
	}

	for _, tt := range testCase {
		repo.Mock.On("GetAccountDetail", tt.Exp.AccountNumber).Return(tt.Exp)
		res, _ := lib.CheckBalance(context.Background(), tt.Exp.AccountNumber)
		assert.Equal(t, tt.Exp, res, "must equal")
	}
}
