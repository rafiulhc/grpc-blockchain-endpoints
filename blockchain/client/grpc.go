package main

import (
	"context"
	"io"
	"log"

	pb "github.com/rafiulhc/grpc-blockchain-endpoints/blockchain/proto"
)



type Block struct {
BlockId int32
Block string
}


	func CallBlock(client pb.GetLatestBlockServiceClient) {
		println("callBlock client called")
		stream, err := client.Block(context.Background(), &pb.GetLatestBlockRequest{})
		if err != nil {
			log.Fatalf("Error while calling Block RPC: %v", err)
		}
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				// we've reached the end of the stream
				break
			}
			if err != nil {
				log.Fatalf("Error while reading stream: %v", err)
			}
			log.Printf("Response from Block: %v", msg.Block)
		}
	}


