package dbo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	db2 "github.com/sentinel-official/explorer/db"
	"github.com/sentinel-official/explorer/types"
)

func BlockInsertOne(ctx context.Context, db *mongo.Database, block types.Block) (interface{}, error) {
	collection := db2.GetBlockCollection(db)

	res, err := collection.InsertOne(ctx, block)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

func BlockFindOne(ctx context.Context, db *mongo.Database, filter interface{}) (*mongo.SingleResult, error) {

	blockCollection := db2.GetBlockCollection(db)

	res := blockCollection.FindOne(ctx, filter)
	if res.Err() != nil {
		return nil, res.Err()
	}

	return res, nil
}

func BlockFindAll(ctx context.Context, db *mongo.Database) (*mongo.Cursor, error) {

	blockCollection := db2.GetBlockCollection(db)

	cursor, err := blockCollection.Find(ctx, bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}

	return cursor, nil
}
