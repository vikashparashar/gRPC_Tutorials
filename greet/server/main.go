package main

import (
	"context"
	"fmt"
	pb "github/vikashparashar/gRPC_Tutorials/greet/proto"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
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
	pb.RegisterGreetServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("server is failed to server to client : %v\n", err)
	}
}
func (s *Server) Single_Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("____Single_Greet_Function Was Invoked At Server____")
	return &pb.GreetResponse{Result: "Hello" + in.Name}, nil
}
func (s *Server) Greet_Many_Times(in *pb.GreetRequest, stream pb.GreetService_Greet_Many_TimesServer) error {
	log.Println("____Greet_Many_Times_Function Was Invoked At Server____")
	for i := 0; i < 5; i++ {
		log.Println("Sending Response To Server !")
		err := stream.Send(&pb.GreetResponse{Result: fmt.Sprintf("Hello Mr. %v , Times : %v", in.Name, i)})
		if err != nil {
			log.Fatalf("Server Streaming Error : %v\n", err)
		}
	}
	return nil
}

func (s *Server) Long_Greet(stream pb.GreetService_Long_GreetServer) error {
	log.Println("____Long_Greet_Function Was Invoked At Server____")
	for {
		var res string
		req, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.GreetResponse{Result: res})
		}
		if err != nil {
			log.Fatalln("Internal Server Error")
		}
		res += fmt.Sprintf("Hello %v , You Are Welcome To gRPC Tutorial", req.Name)
	}
}

// func (s *Server) Greet_Every_One(stream pb.GreetService_Greet_Every_OneServer) error {
// 	log.Println("____Greet_Every_One_Function Was Invoked At Server____")
// }
