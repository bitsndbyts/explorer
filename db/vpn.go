package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetNodeCollection(db *mongo.Database) *mongo.Collection {
	collection := db.Collection("node")

	return collection
}

func GetSubscriptionCollection(db *mongo.Database) *mongo.Collection {
	collection := db.Collection("subscription")

	return collection
}

func GetSessionCollection(db *mongo.Database) *mongo.Collection {
	collection := db.Collection("session")

	return collection
}

func GetProposalCollection(db *mongo.Database) *mongo.Collection {
	collection := db.Collection("proposal	")

	return collection
}

func GetAccountCollection(db *mongo.Database) *mongo.Collection {
	collection := db.Collection("account	")

	return collection
}
