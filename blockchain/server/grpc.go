package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	pb "github.com/rafiulhc/grpc-blockchain-endpoints/blockchain/proto"
)


func (s *Server) Block(req *pb.GetLatestBlockRequest, stream pb.GetLatestBlockService_BlockServer) error{
		println("Block stream called")

		response, err := http.Get("https://rpc.osmosis.zone/abci_info?")

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		for i:=0; i<10; i++{

			res := &pb.GetLatestBlockResponse{
				BlockId: int32(i),
				Block:  string(responseData),
			}
			stream.Send(res)
			time.Sleep(3 * time.Second)
		}
		return nil
	}


