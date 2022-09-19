package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"log"
	"testing"

	pb "github.com/rafiulhc/grpc-blockchain-endpoints/blockchain/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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



func newCallLatestBlock(client pb.GetLatestBlockServiceClient) (result *pb.GetLatestBlockResponse) {

	println("callBlock client called")
	stream, err := client.GetLatestBlock(context.Background(), &pb.GetLatestBlockRequest{})
	if err != nil {
		log.Fatalf("Error while calling Block RPC: %v", err)
	}

		res, err := stream.Recv()

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}


	println("callBlock client finished")
	return res

}

func TestCallLatestBlock(t *testing.T){
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}

	defer conn.Close()

	client := pb.NewGetLatestBlockServiceClient(conn)
	result  := newCallLatestBlock(client)
	println(result.Block)

	responseBlockHashByGRPCRequest := result.Block

	println("Block stream called")

	rpcURL := "https://rpc.osmosis.zone/abci_info?"
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

	responseBlockHashByBlockChainAPIRequest := responseObject.Result.Response.LastBlockAppHash


	if responseBlockHashByGRPCRequest != responseBlockHashByBlockChainAPIRequest {
        t.Error("result doesn't match, got", responseBlockHashByBlockChainAPIRequest)
    }

}