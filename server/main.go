package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	// pb "github.com/gowithvikash/grpc_with_go/greet/proto"
)

var (
	network = "tcp"
	address = "0.0.0.0:50051"
)

type Server struct {
}

func main() {
	fmt.Println("Server Is Trying To Get Ready ! Once Done , Client Can Connect With Server And Can Send Requests To Get Response .")
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatal("Server Listning Error !")
	}
	fmt.Println("Server Is Listning On Address : ", address, " Successfully .")
	s := grpc.NewServer()
	if err = s.Serve(lis); err != nil {
		log.Fatal("Server Serving Failed !")
	}
	fmt.Println("Server Is Ready To Server On Address : ", address)

}

func (s *Server) Simple_Greet() {
	fmt.Println("_____________   Simple_Greet() Function Was Invoked At Server   _____________")
}
func (s *Server) Greet_Many_Times() {
	fmt.Println("___________   Greet_Many_Times() Function Was Invoked At Server   ___________")
}
func (s *Server) Long_Greet() {
	fmt.Println("____________     Long_Greet() Function Was Invoked At Server    _____________")
}
func (s *Server) Greet_Every_One() {
	fmt.Println("____________  Greet_Every_One() Function Was Invoked At Server  _____________")
}
