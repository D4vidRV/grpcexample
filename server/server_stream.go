package main

import (
	"log"
	"time"

	pb "github.com/D4vidRV/grpcexample/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Got request with names: %v", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}

		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}
