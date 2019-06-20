package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTxCollection(db *mongo.Database) *mongo.Collection {
	collection := db.Collection("tx")
	return collection
}
