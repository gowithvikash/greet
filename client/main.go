package main

import (
	"fmt"
	"log"

	pb "github.com/gowithvikash/grpc_with_go/greet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// import (
// )

var (
	address = "localhost:50051"
)

func main() {
	fmt.Println("__Client Is Trying To Connect With Server On Address : ", address)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("__Client Is Failed To Connect With Server !")
	}
	fmt.Println("__Client Is Successfully Connected To Server . ")
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)

	do_Simple_Greet(c)
	// do_Greet_Many_Times(c)
	// do_Long_Greet(c)
	// do_Greet_Every_One(c)
}

func do_Simple_Greet(c pb.GreetServiceClient) {
	fmt.Println("_______   do_Simple_Greet() Function Was Invoked At Client   _______")
}
func do_Greet_Many_Times(c pb.GreetServiceClient) {
	fmt.Println("_______ do_Greet_Many_Times() Function Was Invoked At Client _______")
}
func do_Long_Greet(c pb.GreetServiceClient) {
	fmt.Println("_______    do_Long_Greet() Function Was Invoked At Client    _______")
}
func do_Greet_Every_One(c pb.GreetServiceClient) {
	fmt.Println("_______ do_Greet_Every_One()  Function Was Invoked At Client _______")
}
