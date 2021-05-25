package usecases

import (
	"LinkaAja/src/models"
	"context"
	"errors"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func (u *uc) GetAllData() (interface{}, error) {
	var (
		res       models.AllData
		customers []models.CustomerModel
		accounts  []models.AccountModel
	)

	cursor, err := u.repo.FindAll(bson.M{}, "customer", nil)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		customer := models.CustomerModel{}
		err := cursor.Decode(&customer)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	res.Customers = customers

	cursor, err = u.repo.FindAll(bson.M{}, "account", nil)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		account := models.AccountModel{}
		err := cursor.Decode(&account)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	res.Accounts = accounts

	return res, nil
}

func (u *uc) GetAccountID(urlstring string) (int, error) {
	urlPart := strings.Split(urlstring, "/")
	account_id, err := strconv.Atoi(urlPart[2])
	if err != nil {
		return 0, errors.New("Failed to Get Params")
	}
	return account_id, nil
}
