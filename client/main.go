package main

import "fmt"

// import (
// 	pb "github.com/gowithvikash/grpc_with_go/greet/proto"
// )

var (
	address = "localhost:50051"
)

func main() {
	fmt.Println("Client Is Trying To Connect With Server On Address : ", address)

	do_Simple_Greet()
	do_Greet_Many_Times()
	do_Long_Greet()
	do_Greet_Every_One()
}

func do_Simple_Greet() {
	fmt.Println("-------- do_Simple_Greet() Function Was Invoked At Client --------")
}
func do_Greet_Many_Times() {
	fmt.Println("-------- do_Greet_Many_Times() Function Was Invoked At Client --------")
}
func do_Long_Greet() {
	fmt.Println("-------- do_Long_Greet() Function Was Invoked At Client --------")
}
func do_Greet_Every_One() {
	fmt.Println("-------- do_Greet_Every_One() Function Was Invoked At Client --------")
}
