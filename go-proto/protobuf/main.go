package main

import (
	"fmt"
	//import the generated code
	pb "protobuf/tutorialpb"
	car_pb "protobuf/tutorialpb/car"
	hello_rpc "protobuf/tutorialpb/rpc/hello"

	"google.golang.org/protobuf/proto"
)

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "Roger",
		Email: "test@test.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "13112345678", Type: pb.Person_MOBILE},
			{Number: "01087654321", Type: pb.Person_HOME},
		},
	}

	body, _ := proto.Marshal(p)

	p1 := &pb.Person{}
	_ = proto.Unmarshal(body, p1)

	fmt.Println("Original struct loaded from proto file:", p)
	fmt.Println("Marshalled proto data: ", body)
	fmt.Println("Unmarshalled struct: ", p1)
	fmt.Println("Unmarshalled struct name: ", p1.Name)

	lamborghini := &car_pb.Car{
		Id:     1234,
		Brand:  "Lamborghini",
		Model:  "Aventador",
		Colors: []string{"red", "black", "white"},
	}

	body, _ = proto.Marshal(lamborghini)
	lambo := &car_pb.Car{}

	_ = proto.Unmarshal(body, lambo)

	fmt.Println("Original struct loaded from proto file:", lamborghini)
	fmt.Println("Marshalled proto data: ", body)
	fmt.Println("Unmarshalled struct: ", lambo)
	fmt.Println("Unmarshalled struct brand: ", lambo.GetBrand())
	lambo.Brand = "Ferrari"
	fmt.Println("Unmarshalled struct brand: ", lambo.GetBrand())
	fmt.Println("Original struct loaded from proto file:", lamborghini.GetBrand())

	rpc_hello := &hello_rpc.HelloRequest{
		Greeting: "Hello, World!",
	}

	//call the rpc
	fmt.Println("RPC request: ", rpc_hello.GetGreeting())

}
