package main

import (
	"context"
	"log"
	"net"

	pb "github.com/dimash200101/calculator-gRPC-golang/proto"

)

const (
	protocol = "tcp"
	port     = ":4001"
)

type server struct{}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddReply, error) {
	return &pb.AddReply{N1: in.N1 + in.N2}, nil
}


func main() {
	lis, err := net.Listen(protocol, port)
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
