package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ironman0x7b2/sentinel-sdk/app/hub"

	"github.com/tendermint/tendermint/rpc/client"

	tmTypes "github.com/tendermint/tendermint/types"

	"github.com/ironman0x7b2/explorer/db"
	"github.com/ironman0x7b2/explorer/utils"
)

const (
	rpcServer     = "localhost:26657"
	eventNewBlock = "tm.event='NewBlock'"
	eventTx       = "tm.event='Tx'"
)

func main() {
	fmt.Println("READY")
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

	var cdc = hub.MakeCodec()

	var block_ tmTypes.EventDataNewBlock
	_db := db.DB()

	go func() {
		for e := range block {
			bz, err := json.Marshal(e.Data)
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal(bz, &block_)
			if err != nil {
				panic(err)
			}

			err = utils.ProcessAndInsertBlock(ctx, _db, block_)
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

			err = utils.ProcessAndInsertTx(ctx, cdc, _db, data, e.Tags["action"], e.Tags["sender"])
			if err != nil {
				panic(err)
			}
		}
	}()
	select {}

}
