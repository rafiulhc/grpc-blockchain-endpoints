package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	pb "github.com/rafiulhc/grpc-blockchain-endpoints/blockchain/proto"
)



type Block struct {
BlockId int32
Block string
}

func CallBlock(client pb.GetLatestBlockServiceClient) {
	log.Printf("CallBlock function was invoked with %v", client)
	res, err := client.Block(context.Background(), &pb.GetLatestBlockRequest{})
	if err != nil {
		log.Fatalf("error while calling Block RPC: %v", err)
	}
	log.Printf("Response from Block: %v", res)

	data := Block{
        BlockId: res.BlockId,
		Block: res.Block,
	}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("state.json", file, 0644)
}