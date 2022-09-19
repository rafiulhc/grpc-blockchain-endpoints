package main

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/rafiulhc/grpc-blockchain-endpoints/blockchain/proto"
	"github.com/sirupsen/logrus"
)



type Block struct {
	Height string `json:"height"`
	Hash string   `json:"hash"`
}

func checkFile(filename string) error {
    _, err := os.Stat(filename)
    if os.IsNotExist(err) {
        _, err := os.Create(filename)
        if err != nil {
            return err
        }
    }
    return nil
}



func CallLatestBlock(client pb.GetLatestBlockServiceClient) {

		println("callBlock client called")
		stream, err := client.GetLatestBlock(context.Background(), &pb.GetLatestBlockRequest{})
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

			filename := "state.json"
			err = checkFile(filename)
			if err != nil {
				logrus.Error(err)
			}

			file, err := ioutil.ReadFile(filename)
			if err != nil {
				logrus.Error(err)
			}

			dataBlock := &Block{
				Height: res.BlockId,
				Hash: res.Block,
			}

			data := []Block{}

			// Here the magic happens!
			json.Unmarshal(file, &data)

			data = append(data, *dataBlock)

			// Preparing the data to be marshalled and written.
			dataBytes, err := json.Marshal(data)
			if err != nil {
				logrus.Error(err)
			}

			err = ioutil.WriteFile(filename, dataBytes, 0644)
			if err != nil {
				logrus.Error(err)
			}
			log.Printf("Response from Block: %v", dataBlock)
	}

}