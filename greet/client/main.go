package main

import (
	"context"
	"fmt"
	pb "github/vikashparashar/gRPC_Tutorials/greet/proto"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	address = "localhost:50051" //default port for cliant to connect with grpc
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("client is failed to start a connection with server")
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	do_Single_Greet(c)
	do_Greet_Many_Times(c)
	do_Long_Greet(c)
	// do_Greet_Every_One(c)

}
func do_Single_Greet(c pb.GreetServiceClient) {
	log.Println("____do_Single_Greet_Function Was Invoked By Client____")
	res, err := c.Single_Greet(context.Background(), &pb.GreetRequest{Name: "Vikash Parashar"})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("____do_Single_Greet()_Result: %v\n", res.Result)
}
func do_Greet_Many_Times(c pb.GreetServiceClient) {
	log.Println("____do_Greet_Many_Times_Function Was Invoked By Client____")
	stream, err := c.Greet_Many_Times(context.Background(), &pb.GreetRequest{Name: "Vikash"})
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
		fmt.Printf("____do_Greet_Many_Times()_Result: %v\n", res.Result)
	}

}
func do_Long_Greet(c pb.GreetServiceClient) {
	log.Println("____do_Long_Greet_Function Was Invoked By Client____")
	var req []*pb.GreetRequest = []*pb.GreetRequest{
		{Name: "Vikash"},
		{Name: "Vishnu"},
		{Name: "Vijay"},
		{Name: "Vivaan"},
	}
	stream, err := c.Long_Greet(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range req {
		if err = stream.Send(v); err != nil {
			log.Fatalln(err)
		}
	}
	// if err = stream.CloseSend(); err != nil {
	// 	log.Fatal(err)
	// }
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("____do_Long_Greet()_Result: %v\n", res.Result)
}

// func do_Greet_Every_One(c pb.GreetServiceClient) {
// 	log.Println("____do_Greet_Every_One_Function Was Invoked By Client____")
// }
