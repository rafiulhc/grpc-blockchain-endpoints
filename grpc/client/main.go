package main

import (
	"log"

	pb "github.com/rafiulhc/grpc-blockchain-endpoints/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//
var portAddress string = "localhost:50051"

func main() {
	// connection to the server with insecure credentials for simplicity
	connection, err := grpc.Dial(portAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}

	defer connection.Close()

	// create a client to the server with the connection
	client := pb.NewGetLatestBlockServiceClient(connection)

	// call gRPC client to get latest block
	CallLatestBlock(client)

}