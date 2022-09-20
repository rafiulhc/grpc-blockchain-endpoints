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

var addr string = "0.0.0.0:50051"

// Response is the struct response from the blockchain endpoint
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


// newLatestBlock is a request function that calls the GRPC server and returns the latest block for testing
func newLatestBlock(client pb.GetLatestBlockServiceClient) (*pb.GetLatestBlockResponse) {

	// variable to store the response
	stream, err := client.GetLatestBlock(context.Background(), &pb.GetLatestBlockRequest{})
	if err != nil {
		log.Fatalf("Error while calling Block RPC: %v", err)
	}

	// read the response from the stream
	response, err := stream.Recv()

	if err != nil {
			status.Error(
				codes.Unknown,
				fmt.Sprintf("failed to receive a block response: %v", err),
			)
	}

	return response

}

// TestGetLatestBlock is a test function that tests the GRPC server response with the blockchain API response
func TestCallLatestBlock(t *testing.T){
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		status.Error(
			codes.Internal,
			fmt.Sprintf("failed to connect: %v", err),
		)
	}

	defer conn.Close()

	client := pb.NewGetLatestBlockServiceClient(conn)
	result  := newLatestBlock(client)
	println(result.Block)

	// Call the blockchain endpoint
	rpcURL := "https://rpc.osmosis.zone/abci_info?"
	response, err := http.Get(rpcURL)

	if err != nil {
		status.Error(
			codes.NotFound,
			fmt.Sprintf("Could not get latest block: %v", err),
		)
	}


	body, err := ioutil.ReadAll(response.Body)
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

	responseBlockHashByGRPCRequest := result.Block
	responseBlockHashByBlockChainAPIRequest := responseObject.Result.Response.LastBlockAppHash

	// check if the response from the blockchain API and the response from the GRPC request are the same
	if responseBlockHashByGRPCRequest != responseBlockHashByBlockChainAPIRequest {
        t.Error("result doesn't match, got", responseBlockHashByBlockChainAPIRequest)
    }

}