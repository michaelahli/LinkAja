package models

type TransferBalance struct {
	ToAccountNumber int `json:"to_account_number"`
	Amount          int `json:"amount"`
}
