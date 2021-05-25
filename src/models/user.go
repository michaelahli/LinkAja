package models

type UserModel struct {
	CustomerNumber int    `json:"customer_number" bson:"customer_number"`
	CustomerName   string `json:"customer_name" bson:"customer_name"`
	AccountNumber  int    `json:"account_number" bson:"account_number"`
	Balance        int    `json:"balance" bson:"balance"`
}
