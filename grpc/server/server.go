package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	pb "github.com/rafiulhc/grpc-blockchain-endpoints/grpc/proto"
)

// struct that contains the block height and hash to be written to the json file
type Response struct {
    JSONRPC       string `json:"jsonrpc"`
    ID            int    `json:"id"`
    Result        struct {
		Response   struct {
			Data             string `json:"data"`
			Version          string `json:"version"`
			AppVersion       string `json:"app_version"`
			LastBlockHeight  string `json:"last_block_height"`
			LastBlockAppHash string `json:"last_block_app_hash"`
		  } `json:"response"`
    } `json:"result"`
}

// server call to the osmosis rpc to get the latest block
func (s *Server) GetLatestBlock(req *pb.GetLatestBlockRequest, stream pb.GetLatestBlockService_GetLatestBlockServer) error{
		println("GetLatestBlockService stream called")

		rpcURL := "https://rpc.osmosis.zone/abci_info?"

		// iterate through the rpc url to get the latest block for 5 times
		for i:=0; i<5; i++{
			response, err := http.Get(rpcURL)

			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}

			// read the response body
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			var responseObject Response
			if err := json.Unmarshal(body, &responseObject); err != nil {   // Parse []byte to go struct pointer
				fmt.Println("Can't unmarshal JSON")
			}

			// get the block height and hash from the response body and send it to the client
			result := &pb.GetLatestBlockResponse{
				BlockId: responseObject.Result.Response.LastBlockHeight,
				Block:  responseObject.Result.Response.LastBlockAppHash,
			}
			// send the response to the client as a stream
			stream.Send(result)
			// sleep for 6 seconds as osmosis block execution time is approx. 6 seconds
			time.Sleep(6 * time.Second)
		}

		return nil
	}


