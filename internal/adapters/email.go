package adapters

import "go.mongodb.org/mongo-driver/mongo"

type EmailAdapter struct {
	DB *mongo.Database
}

func NewEmailAdapter(db *mongo.Database) *EmailAdapter{
	return &EmailAdapter{
		DB: db,
	}
}