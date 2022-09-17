package main

import (
	"context"
	"log"

	pb "github.com/rafiulhc/grpc-blockchain-endpoints/blockchain/proto"
)

func (s *Server) Block(ctx context.Context, req *pb.GetLatestBlockRequest) (*pb.GetLatestBlockResponse, error) {
	log.Printf("Block call was invoked with %s", req)
	return &pb.GetLatestBlockResponse{
		BlockId: 12,
		Block:   "RAFIUL",
	}, nil

}