package usecases

import "LinkaAja/src/repositories"

type uc struct {
	repo repositories.Repo
}

type UC interface {
	GetAccountID(urlstring string) (int, error)
	GetAllData() (interface{}, error)
}

func NewUC(repo repositories.Repo) UC {
	return &uc{repo: repo}
}
