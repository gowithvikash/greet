package main

import (
	"context"
	"fmt"
	"io"
	"log"
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
	time.Sleep(1 * time.Second)
	do_Simple_Greet(c)
	time.Sleep(1 * time.Second)
	do_Greet_Many_Times(c)
	time.Sleep(1 * time.Second)
	do_Long_Greet(c)
	time.Sleep(1 * time.Second)
	do_Greet_Every_One(c)
	time.Sleep(1 * time.Second)

}

func do_Simple_Greet(c pb.GreetServiceClient) {
	fmt.Println("____   do_Simple_Greet() Function Was Invoked At Client   ____")
	res, err := c.Simple_Greet(context.Background(), &pb.GreetRequest{Name: "Vikash Parashar"})
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("do_Simple_Greet Result: %v\n", res.Result)
}

//
func do_Greet_Many_Times(c pb.GreetServiceClient) {
	fmt.Println("____ do_Greet_Many_Times() Function Was Invoked At Client ____")

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
		time.Sleep(1 * time.Second)
		fmt.Printf("do_Greet_Many_Times Result: %v\n", res.Result)
	}

}

//
func do_Long_Greet(c pb.GreetServiceClient) {
	fmt.Println("____    do_Long_Greet() Function Was Invoked At Client    ____")

	var reqs = []*pb.GreetRequest{{Name: "Vikash Parashar"}, {Name: "Khushboo Panday"}, {Name: "Niyati"}, {Name: "Ritika"}, {Name: "Rampati Devi"}}

	stream, err := c.Long_Greet(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range reqs {
		fmt.Println("Sending Requests To Server")
		err = stream.Send(v)
		if err != nil {
			log.Fatal("Error While Sending Requests To Server")
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("do_Long_Greet Result: %v\n", res.Result)
}

func do_Greet_Every_One(c pb.GreetServiceClient) {
	fmt.Println("____ do_Greet_Every_One()  Function Was Invoked At Client ____")
	var reqs = []*pb.GreetRequest{{Name: "Vikash Parashar"}, {Name: "Khushboo Panday"}, {Name: "Niyati"}, {Name: "Ritika"}, {Name: "Rampati Devi"}}

	stream, err := c.Long_Greet(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	waitc := make(chan struct{})
	go func() {
		for _, v := range reqs {
			fmt.Println("Sending Requests To Server")
			stream.Send(v)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			res, err := stream.CloseAndRecv() // here it must be recieve only
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("do_Greet_Every_One Result: %v\n", res.Result)

		}
		close(waitc)
	}()
	<-waitc

}
