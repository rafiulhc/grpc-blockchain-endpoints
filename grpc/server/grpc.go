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


func (s *Server) GetLatestBlock(req *pb.GetLatestBlockRequest, stream pb.GetLatestBlockService_GetLatestBlockServer) error{
		println("Block stream called")

		rpcURL := "https://rpc.osmosis.zone/abci_info?"

		for i:=0; i<5; i++{
			response, err := http.Get(rpcURL)

			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}


			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			var responseObject Response

			if err := json.Unmarshal(body, &responseObject); err != nil {   // Parse []byte to go struct pointer
				fmt.Println("Can not unmarshal JSON")
			}

			res := &pb.GetLatestBlockResponse{
				BlockId: responseObject.Result.Response.LastBlockHeight,
				Block:  responseObject.Result.Response.LastBlockAppHash,
			}

			stream.Send(res)

			time.Sleep(6 * time.Second)
		}
		return nil
	}


