package services

import (
	"context"
	"log"
	"time"

	pb "github.com/D4vidRV/grpcexample/proto"
)

func CallSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Printf("Could not send names: %v", err.Error())
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err.Error())
		}
		log.Printf("Send the request with name: %v", name)
		time.Sleep(time.Second * 2)
	}

	res, err := stream.CloseAndRecv()
	log.Println("Client streaming finished")
	if err != nil {
		log.Fatalf("Error while receiving %v", err.Error())
	}
	log.Printf("%v", res.Messages)
}
