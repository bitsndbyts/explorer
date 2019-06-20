package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DB() *mongo.Database {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client.Database("explorer")
}
