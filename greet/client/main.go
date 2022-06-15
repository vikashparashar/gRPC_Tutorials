package main

import (
	pb "github/vikashparashar/gRPC_Tutorials/greet/proto"
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
	c := pb.NewGreetServiceClient(conn)
	do_Single_Greet(c)
	do_Greet_Many_Times(c)
	do_Long_Greet(c)
	do_Greet_Every_One(c)

}
func do_Single_Greet(c pb.GreetServiceClient) {
	log.Println("____do_Single_Greet_Function Was Invoked By Client____")
}
func do_Greet_Many_Times(c pb.GreetServiceClient) {
	log.Println("____do_Greet_Many_Times_Function Was Invoked By Client____")
}
func do_Long_Greet(c pb.GreetServiceClient) {
	log.Println("____do_Long_Greet_Function Was Invoked By Client____")
}
func do_Greet_Every_One(c pb.GreetServiceClient) {
	log.Println("____do_Greet_Every_One_Function Was Invoked By Client____")
}
