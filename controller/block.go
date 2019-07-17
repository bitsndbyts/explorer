package controller

import (
	"context"
	"fmt"

	tmTypes "github.com/tendermint/tendermint/types"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sentinel-official/explorer/dbo"
	"github.com/sentinel-official/explorer/types"
)

func ProcessAndInsertBlock(ctx context.Context, db *mongo.Database, block *tmTypes.Block) error {
	var _block types.Block

	_block.ID = block.LastCommit.BlockID.String()
	_block.Height = block.Height
	_block.DataHash = block.DataHash.String()
	_block.EvidenceHash = block.EvidenceHash.String()
	_block.ValidatorsHash = block.ValidatorsHash.String()
	_block.ProposerAddress = block.ProposerAddress.String()
	_block.NoOfTransactions = block.NumTxs

	id, err := dbo.BlockInsertOne(ctx, db, _block)
	if err != nil {
		return err
	}

	fmt.Println("Inserted block", id)

	return nil
}
