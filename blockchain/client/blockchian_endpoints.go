package main

import (
	"context"
	"log"

	pb "github.com/rafiulhc/grpc-blockchain-endpoints/blockchain/proto"
)

func CallBlock(client pb.BlockServiceClient) {
	log.Printf("doGreet function was invoked with %v", client)
	res, err := client.Block(context.Background(), &pb.BlockRequest{
		Name: "Rafiul",
	})
	if err != nil {
		log.Fatalf("error while calling Block RPC: %v", err)
	}
	log.Printf("Response from Block: %v", res.Message)
}