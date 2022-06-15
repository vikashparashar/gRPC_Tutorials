package main

import (
	pb "github/vikashparashar/gRPC_Tutorials/calculator/proto"
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
	do_Sum(c)
	do_Primes(c)
	do_Average(c)
	do_Max(c)
	do_Sqrt(c)

}

func do_Sum(c pb.CalculatorServiceClient) {
	log.Println("____do_Sum_Function Was Invoked By Client____")
}
func do_Primes(c pb.CalculatorServiceClient) {
	log.Println("____do_Primes_Function Was Invoked By Client____")
}
func do_Average(c pb.CalculatorServiceClient) {
	log.Println("____do_Average_Function Was Invoked By Client____")
}
func do_Max(c pb.CalculatorServiceClient) {
	log.Println("____do_Max_Function Was Invoked By Client____")
}
func do_Sqrt(c pb.CalculatorServiceClient) {
	log.Println("____do_Sqrt_Function Was Invoked By Client____")
}
