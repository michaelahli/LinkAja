package models

type AllData struct {
	Customers []CustomerModel `json:"customers"`
	Accounts  []AccountModel  `json:"accounts"`
}

type BasicResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type GetBalance struct {
	AccountNumber int    `json:"account_number"`
	CustomerName  string `json:"customer_name"`
	Balance       int    `json:"balance"`
}
