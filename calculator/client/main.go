package main

import (
	"context"
	"fmt"
	pb "github/vikashparashar/gRPC_Tutorials/calculator/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

var (
	address = "localhost:50051" //default port for cliant to connect with grpc
)

func main() {
	conn, err := grpc.Dial(address)
	if err != nil {
		log.Fatalln("client is failed to start a connection with server")
	}
	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)
	do_Sum(c, 100, 117)
	do_Primes(c, 128)
	// do_Average(c)
	// do_Max(c)
	// do_Sqrt(c)

}

func do_Sum(c pb.CalculatorServiceClient, n1 int64, n2 int64) {
	var req *pb.SumRequest = &pb.SumRequest{NumberOne: n1, NumberTwo: n2}
	log.Println("____do_Sum_Function Was Invoked By Client____")
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("_____do_Sum Function () Result: %v\n", res.Result)
}

func do_Primes(c pb.CalculatorServiceClient, n int64) {
	log.Println("____do_Primes_Function Was Invoked By Client____")
	var req *pb.PrimeRequest = &pb.PrimeRequest{Number: n}
	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("_____do_Primes Function () Result: %v\n", res.Result)
	}
}

func do_Average(c pb.CalculatorServiceClient) {
	log.Println("____do_Average_Function Was Invoked By Client____")
	var requests []*pb.AverageRequest = []*pb.AverageRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
		{Number: 5},
		{Number: 6},
		{Number: 7},
		{Number: 8},
		{Number: 9},
		{Number: 10},
	}
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range requests {
		stream.Send(v)
	}
	// err = stream.CloseSend()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("____do_Average_Function Result: %v\n", res.Result)

}

// func do_Max(c pb.CalculatorServiceClient) {
// 	log.Println("____do_Max_Function Was Invoked By Client____")
// }
// func do_Sqrt(c pb.CalculatorServiceClient) {
// 	log.Println("____do_Sqrt_Function Was Invoked By Client____")
// }
