package main

import (
	"context"

	pb "haseeb.khan/grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello Haseeb",
	}, nil
}
