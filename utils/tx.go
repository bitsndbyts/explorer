package utils

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	tmTypes "github.com/tendermint/tendermint/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	db2 "github.com/ironman0x7b2/explorer/db"
	"github.com/ironman0x7b2/explorer/types"
)

func ProcessAndInsertTx(ctx context.Context, cdc *codec.Codec, db *mongo.Database,
	data tmTypes.EventDataTx, _type, sender string) error {
	var t types.Tx
	txCollection := db2.GetTxCollection(db)

	blockCollection := db2.GetBlockCollection(db)
	var block types.Block

	err := blockCollection.FindOne(ctx, bson.D{{
		"height", data.Height,
	}}).Decode(&block)

	if err != nil {
		return err
	}

	t.BlockID = block.ID
	t.Height = data.Height
	//t.Hash =
	t.Type = _type
	t.From = sender
	t.GasUsed = data.Result.GasUsed
	t.GasWanted = data.Result.GasWanted

	if data.Result.Code == uint32(0){
		t.Status = true
	}

	var ttt auth.StdTx
	err = cdc.UnmarshalBinaryLengthPrefixed(data.TxResult.Tx, &ttt)
	if err != nil {
		return err
	}

	t.Fee = ttt.Fee.Amount.String()

	//t.Msg = ttt.GetMsgs() TODO: Need to store msg

	res, err := txCollection.InsertOne(ctx, t)
	if err != nil {
		return err
	}
	fmt.Println("tx inserted", res.InsertedID)

	return nil
}
