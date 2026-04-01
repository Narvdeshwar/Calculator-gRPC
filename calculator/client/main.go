package main

import (
	"context"
	pb "grpc-calculator/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error at client", err)
	}

	defer conn.Close()

	client := pb.NewCalculatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.Multiply(ctx, &pb.CalcRequest{
		A: 20,
		B: 15,
	})
	if err != nil {
		log.Fatal("Error")
	}
	log.Println("Result:", res.Result)

}
