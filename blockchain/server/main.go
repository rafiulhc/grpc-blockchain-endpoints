package main

import (
	"log"
	"net"

	pb "github.com/rafiulhc/grpc-blockchain-endpoints/blockchain/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50051"

type Server struct{
	pb.GetLatestBlockServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	log.Printf("Server is listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGetLatestBlockServiceServer(s, &Server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}