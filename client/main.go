package main

import (
	"fmt"
	"log"

	"github.com/D4vidRV/grpcexample/client/services"
	pb "github.com/D4vidRV/grpcexample/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect %v", err.Error())
	}

	// Close connection finally
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{},
	}

	for i := 0; i < 10; i++ {
		names.Names = append(names.Names, fmt.Sprint(i+1))
	}

	// services.CallSayHello(client)
	// services.CallSayHelloServerStream(client, names)
	// services.CallSayHelloClientStream(client, names)
	services.CallHelloBidirectionalStream(client, names)
}
