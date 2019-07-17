package dbo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	db2 "github.com/sentinel-official/explorer/db"
	"github.com/sentinel-official/explorer/types"
)

func NodeInsertOne(ctx context.Context, db *mongo.Database, node types.Node) (interface{}, error) {
	collection := db2.GetNodeCollection(db)

	res, err := collection.InsertOne(ctx, node)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

func NodeFindOne(ctx context.Context, db *mongo.Database, filter interface{}) (*mongo.SingleResult, error) {
	collection := db2.GetNodeCollection(db)

	res := collection.FindOne(ctx, filter)
	if res.Err() != nil {
		return nil, res.Err()
	}

	return res, nil
}

func NodeFindAll(ctx context.Context, db *mongo.Database) (*mongo.Cursor, error) {
	collection := db2.GetNodeCollection(db)

	cursor, err := collection.Find(ctx, bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}

	return cursor, nil
}
