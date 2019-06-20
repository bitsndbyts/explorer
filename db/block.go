package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBlockCollection(db *mongo.Database) *mongo.Collection {
	collection := db.Collection("block")
	return collection
}
