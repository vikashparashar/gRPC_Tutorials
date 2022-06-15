package main

import (
	"context"
	pb "github/vikashparashar/gRPC_Tutorials/calculator/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.CalculatorServiceServer
}

var (
	network = "tcp"
	address = "0.0.0.0:50051" //default port for grpc
)

func main() {
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("server is failed to listen : %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("server is failed to server to client : %v\n", err)
	}
}

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Println("____Sum_Function Was Invoked At Server____")
}
func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Println("____Primes_Function Was Invoked At Server____")
}
func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("____Average_Function Was Invoked At Server____")
}
func (s *Server) Maximum(stream pb.CalculatorService_MaximumServer) error {
	log.Println("____Maximum_Function Was Invoked At Server____")
}
func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Println("____Sqrt_Function Was Invoked At Server____")
}
