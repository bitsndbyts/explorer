package dbo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	db2 "github.com/sentinel-official/explorer/db"
	"github.com/sentinel-official/explorer/types"
)

func AccountInsertOne(ctx context.Context, db *mongo.Database, account types.Account) (interface{}, error) {
	collection := db2.GetAccountCollection(db)

	res, err := collection.InsertOne(ctx, account)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

func AccountFindOne(ctx context.Context, db *mongo.Database, filter interface{}) (*mongo.SingleResult, error) {
	collection := db2.GetAccountCollection(db)

	res := collection.FindOne(ctx, filter)
	if res.Err() != nil {
		return nil, res.Err()
	}

	return res, nil
}

func AccountFindAll(ctx context.Context, db *mongo.Database) (*mongo.Cursor, error) {
	collection := db2.GetAccountCollection(db)

	cursor, err := collection.Find(ctx, bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}

	return cursor, nil
}
