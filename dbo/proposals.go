package dbo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	db2 "github.com/sentinel-official/explorer/db"
	"github.com/sentinel-official/explorer/types"
)

func ProposalInsertOne(ctx context.Context, db *mongo.Database, proposal types.Proposals) (interface{}, error) {
	collection := db2.GetProposalCollection(db)

	res, err := collection.InsertOne(ctx, proposal)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

func ProposalFindOne(ctx context.Context, db *mongo.Database, filter interface{}) (*mongo.SingleResult, error) {
	collection := db2.GetProposalCollection(db)

	res := collection.FindOne(ctx, filter)
	if res.Err() != nil {
		return nil, res.Err()
	}

	return res, nil
}

func ProposalFindAll(ctx context.Context, db *mongo.Database) (*mongo.Cursor, error) {
	collection := db2.GetProposalCollection(db)

	cursor, err := collection.Find(ctx, bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}

	return cursor, nil
}
