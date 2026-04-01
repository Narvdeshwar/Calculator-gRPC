package main

import (
	"context"
	"fmt"
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

	var operation string
	var num1, num2 int32

	fmt.Println("=== gRPC Calculator ===")
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	fmt.Print("Enter the operation (add / multiply): ")
	fmt.Scan(&operation)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var result int32

	if operation == "add" {
		res, err := client.Add(ctx, &pb.CalcRequest{
			A: num1,
			B: num2,
		})
		if err != nil {
			log.Fatal("Error performing the addition:", err)
		}
		result = res.Result

	} else if operation == "multiply" {
		res, err := client.Multiply(ctx, &pb.CalcRequest{
			A: num1,
			B: num2,
		})
		if err != nil {
			log.Fatal("Error performing multiplication:", err)
		}
		result = res.Result

	} else {
		log.Fatal("Galat operation type kiya! Kripya 'add' ya 'multiply' likhein.")
	}

	log.Println("Final Result:", result)
}
