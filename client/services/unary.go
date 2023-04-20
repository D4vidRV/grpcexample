package services

import (
	"context"
	"log"
	"time"

	pb "github.com/D4vidRV/grpcexample/proto"
)

func CallSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err.Error())
	}
	log.Printf("%v", resp.Message)
}
