package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "haseeb.khan/grpc/proto"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Haseeb", "Sabari", "Kapil", "Batool"},
	}
	//? Command to generate proto files
	//TODO: protoc --go_out=. --go-grpc_out=. proto/greet.proto

	//? Unary API call
	// callSayHello(client)

	//? Server side streaming
	//callSayHelloServerStreaming(client, names)

	//? Client side streaming
	// callSayHelloClientStream(client, names)

	//? BiDirectional Streaming
	callHelloBidirectionalStream(client, names)
}
