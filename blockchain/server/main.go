package main

import (
	"log"
	"net"

	pb "github.com/rafiulhc/grpc-blockchain-endpoints/blockchain/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct{
	pb.BlockServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	log.Printf("Server is listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterBlockServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}