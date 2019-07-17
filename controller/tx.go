package controller

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	tmTypes "github.com/tendermint/tendermint/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sentinel-official/explorer/dbo"
	"github.com/sentinel-official/explorer/types"
)

func ProcessAndInsertTx(ctx context.Context, cdc *codec.Codec, db *mongo.Database,
	data tmTypes.EventDataTx, _type, sender string) error {
	var tx types.Tx

	var block types.Block
	filter := bson.D{{"height", data.Height}}
	res, err := dbo.BlockFindOne(ctx, db, filter)
	if err != nil {
		return err
	}

	err = res.Decode(&block)
	if err != nil {
		return err
	}

	tx.BlockID = block.ID
	tx.Height = data.Height
	tx.Type = _type
	tx.From = sender
	tx.GasUsed = data.Result.GasUsed
	tx.GasWanted = data.Result.GasWanted

	if data.Result.Code == uint32(0) {
		tx.Status = true
	}

	//
	//var stdTx auth.StdTx
	//err = cdc.UnmarshalBinaryLengthPrefixed(data.TxResult.Tx, &stdTx)
	//if err != nil {
	//	return err
	//}
	//
	//tx.Fee = stdTx.Fee.Amount.String()

	//TODO : Need to store Tx hash
	//tx.Msg = stdTx.GetMsgs() TODO: Need to store msg

	id, err := dbo.TxInsertOne(ctx, db, tx)
	if err != nil {
		return err
	}
	fmt.Println("tx inserted", id)

	return nil
}
