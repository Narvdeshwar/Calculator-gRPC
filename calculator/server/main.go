package main

import (
	"context"
	pb "grpc-calculator/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, req *pb.CalcRequest) (*pb.CalcResponse, error) {
	result := req.A + req.B
	return &pb.CalcResponse{
		Result: result,
	}, nil
}

func (s *server) Multiply(ctx context.Context, req *pb.CalcRequest) (*pb.CalcResponse, error) {
	result := req.A * req.B
	return &pb.CalcResponse{
		Result: result,
	}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("err", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServer(grpcServer, &server{})
	log.Println("Server started on port 50051")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("err", err)
	}
}
