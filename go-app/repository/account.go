package repository

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/marioheryanto/linkaja/go-app/helper"
	"github.com/marioheryanto/linkaja/go-app/model"
)

type AccountRepository struct {
	dbClient *sqlx.DB
}

type AccountRepositoryInterface interface {
	GetAccountDetail(ctx context.Context, accNumber string) (model.Account, error)
	TransferBalance(ctx context.Context, params model.TransferParams) error
}

func NewAccountRepository(dbClient *sqlx.DB) AccountRepositoryInterface {
	return &AccountRepository{
		dbClient: dbClient,
	}
}

func (r *AccountRepository) GetAccountDetail(ctx context.Context, accNumber string) (model.Account, error) {
	var acc model.Account

	query := `
		SELECT 
		  a.account_number,
		  c.name, 
		  a.balance 
		FROM 
		  accounts a 
		  JOIN customers c ON c.customer_number = a.customer_number 
		WHERE 
		  account_number = ?
  `

	err := r.dbClient.QueryRowContext(ctx, query, accNumber).Scan(&acc.AccountNumber, &acc.CustomerName, &acc.Balance)
	if err != nil && err != sql.ErrNoRows {
		return acc, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	if err == sql.ErrNoRows {
		return acc, helper.NewServiceError(http.StatusBadRequest, "account tidak ditemukan")
	}

	return acc, nil
}

func (r *AccountRepository) TransferBalance(ctx context.Context, params model.TransferParams) error {
	// start trx
	tx, err := r.dbClient.BeginTx(ctx, nil)
	if err != nil {
		return helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	// lock rows
	query := `SELECT account_number, balance FROM accounts where account_number IN(?,?) FOR UPDATE`
	rows, err := tx.QueryContext(ctx, query, params.FromAccountNumber, params.ToAccountNumber)
	if err != nil {
		tx.Rollback()
		return helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	// validate data
	defer rows.Close()
	for rows.Next() {
		account := model.Account{}

		err := rows.Scan(&account.AccountNumber, &account.Balance)
		if err != nil {
			tx.Rollback()
			return helper.NewServiceError(http.StatusInternalServerError, err.Error())
		}

		if account.AccountNumber == params.FromAccountNumber && account.Balance < params.Amount {
			return helper.NewServiceError(http.StatusBadRequest, "saldo tidak cukup")
		}
	}

	if err := rows.Err(); err != nil {
		return helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	// update balance
	query = `UPDATE accounts SET balance = balance - ? WHERE account_number = ?`
	_, err = tx.ExecContext(ctx, query, params.Amount, params.FromAccountNumber)
	if err != nil {
		tx.Rollback()
		return helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	query = `UPDATE accounts SET balance = balance + ? WHERE account_number = ?`
	_, err = tx.ExecContext(ctx, query, params.Amount, params.ToAccountNumber)
	if err != nil {
		tx.Rollback()
		return helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	return tx.Commit()
}
