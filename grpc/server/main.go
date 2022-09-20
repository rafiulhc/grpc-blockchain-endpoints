package main

import (
	"log"
	"net"

	pb "github.com/rafiulhc/grpc-blockchain-endpoints/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var portAddress string = "0.0.0.0:50051"

// struct type of the server that implements the GetLatestBlockServiceServer interface
type Server struct{
	pb.GetLatestBlockServiceServer
}

func main() {
	// listener on the port
	listen, err := net.Listen("tcp", portAddress)

	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	log.Printf("Server is listening on %s\n", portAddress)

	// gRPC server
	server := grpc.NewServer()
	// register the server to the GetLatestBlockServiceServer interface
	pb.RegisterGetLatestBlockServiceServer(server, &Server{})
	// register reflection service on gRPC server calling from CLI by grpcurl
	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}