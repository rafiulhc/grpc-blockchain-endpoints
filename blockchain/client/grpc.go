package main

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"

	pb "github.com/rafiulhc/grpc-blockchain-endpoints/blockchain/proto"
)



type Block struct {
	BlockId string
	Block string
}


	func CallBlock(client pb.GetLatestBlockServiceClient) {
		println("callBlock client called")
		stream, err := client.Block(context.Background(), &pb.GetLatestBlockRequest{})
		if err != nil {
			log.Fatalf("Error while calling Block RPC: %v", err)
		}
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				// we've reached the end of the stream
				break
			}
			if err != nil {
				log.Fatalf("Error while reading stream: %v", err)
			}

			data := &Block{
				BlockId: res.BlockId,
				Block: res.Block,
			}

			file, _ := json.MarshalIndent(data, "", " ")

			_ = ioutil.WriteFile("state.json", file, 0644)
			log.Printf("Response from Block: %v", data)
		}
	}


