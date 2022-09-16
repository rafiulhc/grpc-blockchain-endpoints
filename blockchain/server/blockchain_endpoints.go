package main

import (
	"context"
	"log"

	pb "github.com/rafiulhc/grpc-blockchain-endpoints/blockchain/proto"
)

func (s *Server) Block(ctx context.Context, req *pb.BlockRequest) (*pb.BlockResponse, error) {
	log.Printf("Greet function was invoked with %s", req)
	return &pb.BlockResponse{
		Message: "Hello " + req.Name,
	}, nil

}