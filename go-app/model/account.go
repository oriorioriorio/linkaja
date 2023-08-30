package model

type TransferParams struct {
	FromAccountNumber string
	ToAccountNumber   string  `json:"to_account_number"`
	Amount            float64 `json:"amount"`
}

type Account struct {
	AccountNumber string  `json:"account_number"`
	CustomerName  string  `json:"customer_name"`
	Balance       float64 `json:"balance"`
}
