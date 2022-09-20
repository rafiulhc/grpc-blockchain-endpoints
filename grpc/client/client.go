package main

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/rafiulhc/grpc-blockchain-endpoints/grpc/proto"
	"github.com/sirupsen/logrus"
)

// struct that contains the block height and hash to be written to the json file
type Block struct {
	Height string `json:"height"`
	Hash string   `json:"hash"`
}

// helper function to check if the file exists for appending the new block data
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


// client call to the grpc server to get the latest block
func CallLatestBlock(client pb.GetLatestBlockServiceClient) {

		println("GetLatestBlock client called")
		stream, err := client.GetLatestBlock(context.Background(), &pb.GetLatestBlockRequest{})
		if err != nil {
			log.Fatalf("Error while calling RPC: %v", err)
		}

		// for loop to iterate over response from the grpc server streaming
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				// we've reached the end of the stream
				break
			}
			if err != nil {
				log.Fatalf("Error while reading stream: %v", err)
			}

			filename := "state.json"
			// check if the file exists by calling the helper function
			err = checkFile(filename)
			if err != nil {
				logrus.Error(err)
			}

			file, err := ioutil.ReadFile(filename)
			if err != nil {
				logrus.Error(err)
			}

			dataBlock := &Block{
				Height: response.BlockId,
				Hash: response.Block,
			}

			//
			blockArray := []Block{}

			// unmarshal the json file to get the previous block height and hash
			json.Unmarshal(file, &blockArray)
			// append the new block height and hash to the json file
			blockArray = append(blockArray, *dataBlock)

			// Preparing the data to be marshalled and written.
			dataBytes, err := json.Marshal(blockArray)
			if err != nil {
				logrus.Error(err)
			}

			// write the all existing blocks height and hash to the json file
			err = ioutil.WriteFile(filename, dataBytes, 0644)
			if err != nil {
				logrus.Error(err)
			}
			println("Response from Block: %v", "BlockHash:", response.Block, "Block height:", response.BlockId)
	}

}