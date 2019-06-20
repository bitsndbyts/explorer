package utils

import (
	"context"
	"fmt"

	"github.com/tendermint/tendermint/libs/bech32"
	"go.mongodb.org/mongo-driver/mongo"
	tmTypes "github.com/tendermint/tendermint/types"

	db2 "github.com/ironman0x7b2/explorer/db"
	"github.com/ironman0x7b2/explorer/types"
)

func ProcessAndInsertBlock(ctx context.Context, db *mongo.Database, block tmTypes.EventDataNewBlock) error {
	var b types.Block
	blockCollection := db2.GetBlockCollection(db)

	b.ID = block.Block.LastCommit.BlockID.String()
	b.Height = block.Block.Height
	b.DataHash = block.Block.DataHash.String()
	b.EvidenceHash = block.Block.EvidenceHash.String()
	b.ValidatorsHash = block.Block.ValidatorsHash.String()
	addr, err := bech32.ConvertAndEncode("sent", block.Block.ProposerAddress)
	if err != nil {
		return err
	}
	b.ProposerAddress = addr
	res, err := blockCollection.InsertOne(ctx, b)
	if err != nil {
		return err
	}
	fmt.Println("Inserted block", res.InsertedID)

	return nil
}