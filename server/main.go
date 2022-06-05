package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

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
	time.Sleep(1 * time.Second)
	fmt.Println("\n__Server Is Trying To Get Ready ! Once Done , Client Can Connect With Server And Can Send Requests To Get Response .")
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatal("\n__Server Listning Error !")
	}
	time.Sleep(1 * time.Second)
	fmt.Println("\n__Server Is Listning On Address : ", address, " Successfully .")
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatal("\n__Server Serving Failed !")
	}
	time.Sleep(1 * time.Second)
	fmt.Println("\n__Server Is Ready To Server On Address : ", address)

}

func (s *Server) Simple_Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Println("\n_____   Simple_Greet() Function Was Invoked At Server   _____")
	time.Sleep(1 * time.Second)
	return &pb.GreetResponse{Result: "Hello , " + in.Name + " ! "}, nil
}

func (s *Server) Greet_Many_Times(in *pb.GreetRequest, stream pb.GreetService_Greet_Many_TimesServer) error {
	fmt.Println("\n_____   Greet_Many_Times() Function Was Invoked At Server   _____")
	for i := 1; i < 5; i++ {
		time.Sleep(1 * time.Second)
		if err := stream.Send(&pb.GreetResponse{
			Result: "Hello " + in.Name,
		}); err != nil {
			log.Fatal(err)
		}
	}
	return nil

}

func (s *Server) Long_Greet(stream pb.GreetService_Long_GreetServer) error {
	fmt.Println("\n_____     Long_Greet() Function Was Invoked At Server    _____")
	var res = ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			time.Sleep(1 * time.Second)
			return stream.SendAndClose(&pb.GreetResponse{Result: res})
		}
		if err != nil {
			log.Fatal(err)
		}
		res += fmt.Sprintf("\nHello , %s !\n", req.Name)

	}
}

func (s *Server) Greet_Every_One(stream pb.GreetService_Greet_Every_OneServer) error {
	fmt.Println("_____  Greet_Every_One() Function Was Invoked At Server  _____")
	for {
		req, err := stream.Recv()
		fmt.Println("Receiving Requests From Client")
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatal(err)
		}

		err = stream.Send(&pb.GreetResponse{Result: "Hello" + req.Name})
		if err != nil {
			log.Fatal(err)
		}
	}
}
