package main

import (
	"context"
	"encoding/json"

	"github.com/sentinel-official/hub/app"
	"github.com/tendermint/tendermint/rpc/client"
	tmTypes "github.com/tendermint/tendermint/types"

	"github.com/sentinel-official/explorer/controller"
	"github.com/sentinel-official/explorer/db"
)

const (
	rpcServer     = "localhost:26657"
	eventNewBlock = "tm.event='NewBlock'"
	eventTx       = "tm.event='Tx'"
)

func main() {
	rpcClient := client.NewHTTP(rpcServer, "/websocket")
	err := rpcClient.OnStart()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	block, err := rpcClient.Subscribe(ctx, "subscribe", eventNewBlock)
	if err != nil {
		panic(err)
	}

	cdc := app.MakeCodec()
	_block := tmTypes.EventDataNewBlock{}
	_db := db.DB()

	go func() {
		for e := range block {
			bz, err := json.Marshal(e.Data)
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal(bz, &_block)
			if err != nil {
				panic(err)
			}

			err = controller.ProcessAndInsertBlock(ctx, _db, _block.Block)
			if err != nil {
				panic(err)
			}
		}
	}()

	txData, err := rpcClient.Subscribe(ctx, "subscribe", eventTx)
	if err != nil {
		panic(err)
	}

	var data tmTypes.EventDataTx

	go func() {
		for e := range txData {
			bz, err := json.Marshal(e.Data)
			if err != nil {
				panic(err)
			}

			err = json.Unmarshal(bz, &data)
			if err != nil {
				panic(err)
			}

			err = controller.ProcessAndInsertTx(ctx, cdc, _db, data, e.Tags["action"], e.Tags["sender"])
			if err != nil {
				panic(err)
			}
		}
	}()
	select {}

}
