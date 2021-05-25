package repositories

import "go.mongodb.org/mongo-driver/mongo"

type repo struct {
	db *mongo.Database
}

type Repo interface {
	SampleFunc() error
}

func NewRepo(db *mongo.Database) Repo {
	return &repo{db: db}
}
