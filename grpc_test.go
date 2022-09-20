package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"log"
	"testing"

	pb "github.com/rafiulhc/grpc-blockchain-endpoints/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure" // for simplicity, use insecure.NewCredentials() as a parameter in grpc.Dial(), for production use SSL cert
	"google.golang.org/grpc/status"               // for error handling status
)

var portAddress string = "0.0.0.0:50051"

// Response struct to store response from the blockchain endpoint
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


// request function that calls the GRPC server and returns the latest block for testing
func newLatestBlock(client pb.GetLatestBlockServiceClient) (*pb.GetLatestBlockResponse) {

	// variable to store the response
	stream, err := client.GetLatestBlock(context.Background(), &pb.GetLatestBlockRequest{})
	if err != nil {
		log.Fatalf("Error while calling Block RPC: %v", err)
	}

	// response from the grpc server streaming
	response, err := stream.Recv()

	if err != nil {
			status.Error(
				codes.Unknown,
				fmt.Sprintf("failed to receive a block response: %v", err),
			)
	}

	return response

}

// tester function, compare the GRPC server response with the osmosis blockchain endpoint response
func TestCallLatestBlock(t *testing.T){
	connection, err := grpc.Dial(portAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		status.Error(
			codes.Internal,
			fmt.Sprintf("failed to connect: %v", err),
		)
	}

	defer connection.Close()

	client := pb.NewGetLatestBlockServiceClient(connection)
	resultFromGRPC  := newLatestBlock(client)
	println("Block hash:", resultFromGRPC.Block)

	// Call the osmosis/blockchain endpoint
	rpcURL := "https://rpc.osmosis.zone/abci_info?"
	responseFromBlockchain, err := http.Get(rpcURL)

	if err != nil {
		status.Error(
			codes.NotFound,
			fmt.Sprintf("Could not get latest block: %v", err),
		)
	}


	body, err := ioutil.ReadAll(responseFromBlockchain.Body)
	if err != nil {
		status.Error(
			codes.Internal,
			fmt.Sprintf("Could not read response body: %v", err),
		)
	}

	// Unmarshal the response body and store it in the Response struct
	var responseObject Response

	if err := json.Unmarshal(body, &responseObject); err != nil {   // Parsed []byte to go struct pointer
		status.Error(
			codes.Internal,
			fmt.Sprintf("Could not unmarshal response body: %v", err),
		)
	}

	responseBlockHashByGRPCRequest := resultFromGRPC.Block
	responseBlockHashByBlockChainAPIRequest := responseObject.Result.Response.LastBlockAppHash

	// check if the response from the blockchain API and the response from the GRPC request are the same
	if responseBlockHashByGRPCRequest != responseBlockHashByBlockChainAPIRequest {
        t.Error("result doesn't match, got", responseBlockHashByBlockChainAPIRequest)
    }

}