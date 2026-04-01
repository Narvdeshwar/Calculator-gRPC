package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-calculator/proto"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error at client", err)
	}

	defer conn.Close()

	client := pb.NewCalculatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.Add(ctx, &pb.CalcRequest{
		A: 20,
		B: 15,
	})
	if err != nil {
		log.Fatal("Error")
	}
	log.Println("Result:", res.Result)

}
