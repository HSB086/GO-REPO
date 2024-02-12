package main

import (
	"io"
	"log"

	pb "haseeb.khan/grpc/proto"
)

func (s *helloServer) SayHelloBiDiractionalStreaming(stream pb.GreetService_SayHelloBiDiractionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name + " from Server",
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
