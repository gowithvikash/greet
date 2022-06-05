package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/gowithvikash/grpc_with_go/greet/proto"
	"google.golang.org/grpc"
)

var (
	network = "tcp"
	address = "0.0.0.0:50051"
)

type Server struct {
	pb.GreetServiceServer
}

func main() {
	fmt.Println("__Server Is Trying To Get Ready ! Once Done , Client Can Connect With Server And Can Send Requests To Get Response .")
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatal("__Server Listning Error !")
	}
	fmt.Println("__Server Is Listning On Address : ", address, " Successfully .")
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatal("__Server Serving Failed !")
	}
	fmt.Println("__Server Is Ready To Server On Address : ", address)

}

func (s *Server) Simple_Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Println("_____________   Simple_Greet() Function Was Invoked At Server   _____________")
	return &pb.GreetResponse{Result: "Hello , " + in.Name + " ! "}, nil
}

func (s *Server) Greet_Many_Times(in *pb.GreetRequest, stream pb.GreetService_Greet_Many_TimesServer) error {
	fmt.Println("___________   Greet_Many_Times() Function Was Invoked At Server   ___________")

}

func (s *Server) Long_Greet(stream pb.GreetService_Long_GreetServer) error {
	fmt.Println("____________     Long_Greet() Function Was Invoked At Server    _____________")
}
func (s *Server) Greet_Every_One(stream pb.GreetService_Greet_Every_OneServer) error {
	fmt.Println("____________  Greet_Every_One() Function Was Invoked At Server  _____________")
}
