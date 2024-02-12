package main

import (
	"context"
	"log"
	"time"

	pb "haseeb.khan/grpc/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client Streaming Started...")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sendig %v", err)
		}
		log.Printf("Sent request with names: %v", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client Streaming Finished.")
	if err != nil {
		log.Fatalf("Error While recieving %v", err)
	}
	log.Printf("%v", res.Messages)
}
