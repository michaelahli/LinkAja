package usecases

import (
	"LinkaAja/src/models"
	"LinkaAja/src/repositories"
)

type uc struct {
	repo repositories.Repo
}

type UC interface {
	GetAccountID(urlstring string) (int, error)
	GetAllData() (interface{}, error)
	GetBalanceInfo(account_number int) (*models.GetBalance, error)
}

func NewUC(repo repositories.Repo) UC {
	return &uc{repo: repo}
}
