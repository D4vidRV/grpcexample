package services

import (
	"context"
	"io"
	"log"

	pb "github.com/D4vidRV/grpcexample/proto"
)

func CallSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names: %v", err.Error())
	}

	for {
		message, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while streaming %v", err.Error())
		}

		log.Println(message)
	}
	log.Println("Streaming finished")
}
