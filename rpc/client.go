package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// Tạo Client kết nối đến địa chỉ localhost:1234
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	var reply string
	// Thực hiện lời gọi RPC với service là Greeting và tên method là Hello
	err = client.Call("Greeting.Hello", "Simon", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
