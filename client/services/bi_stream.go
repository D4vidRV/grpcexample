package services

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/D4vidRV/grpcexample/proto"
)

func CallHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: , %v", err.Error())
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while streaming %v", err.Error())
			}

			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err.Error())
		}
		time.Sleep(time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional streaming finished")
}
