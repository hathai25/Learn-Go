package main

import (
	"fmt"
	//import the generated code
	pb "protobuf/tutorialpb"

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
}
