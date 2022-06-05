package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

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
	time.Sleep(1 * time.Second)
	fmt.Println("__Client Is Trying To Connect With Server On Address : ", address)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("__Client Is Failed To Connect With Server !")
	}
	time.Sleep(1 * time.Second)
	fmt.Println("__Client Is Successfully Connected To Server . ")
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	time.Sleep(5 * time.Second)
	do_Simple_Greet(c)
	time.Sleep(5 * time.Second)
	do_Greet_Many_Times(c)
	time.Sleep(5 * time.Second)
	do_Long_Greet(c)
	time.Sleep(5 * time.Second)
	// do_Greet_Every_One(c)
	// time.Sleep(5 * time.Second)
	os.Exit(1)
}

func do_Simple_Greet(c pb.GreetServiceClient) {
	fmt.Println("_______   do_Simple_Greet() Function Was Invoked At Client   _______")
	res, err := c.Simple_Greet(context.Background(), &pb.GreetRequest{Name: "Vikash Parashar"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("do_Simple_Greet Result: %v\n", res.Result)
}

//
func do_Greet_Many_Times(c pb.GreetServiceClient) {
	fmt.Println("_______ do_Greet_Many_Times() Function Was Invoked At Client _______")

	stream, err := c.Greet_Many_Times(context.Background(), &pb.GreetRequest{Name: "Learn To Code"})

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
		fmt.Printf("do_Greet_Many_Times Result: %v\n", res.Result)
	}

}

//
func do_Long_Greet(c pb.GreetServiceClient) {
	fmt.Println("_______    do_Long_Greet() Function Was Invoked At Client    _______")

	var reqs = []*pb.GreetRequest{{Name: "Vikash Parashar"}, {Name: "Khushboo Panday"}, {Name: "Niyati"}, {Name: "Ritika"}, {Name: "Rampati Devi"}}

	stream, err := c.Long_Greet(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range reqs {
		stream.Send(v)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("do_Long_Greet Result: %v\n", res.Result)
}

//
// func do_Greet_Every_One(c pb.GreetServiceClient) {
// 	fmt.Println("_______ do_Greet_Every_One()  Function Was Invoked At Client _______")
// 	var reqs = []*pb.GreetRequest{{Name: "Vikash Parashar"}, {Name: "Khushboo Panday"}, {Name: "Niyati"}, {Name: "Ritika"}, {Name: "Rampati Devi"}}

// 	stream, err := c.Long_Greet(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, v := range reqs {
// 		stream.Send(v)
// 	}

// }
