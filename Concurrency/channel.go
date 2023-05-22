// Go program to illustrate
// how to create a channel
package main

import (
	"fmt"
	"time"
)

func routine1(channel1 chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Send to channel 1")
	channel1 <- 1
}

func routine2(channel2 chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Send to channel 2")
	channel2 <- 2
}

func main() {

	// Creating a channel using var keyword
	var channel1 chan int

	// Creating a channel using make() function
	channel2 := make(chan int)

	// Calling goroutine
	go routine1(channel1)
	go routine2(channel2)

	select {
	case receive1 := <-channel1:
		fmt.Println("Receive from channel 1")
		fmt.Println(receive1)
	case receive2 := <-channel2:
		fmt.Println("Receive from channel 2")
		fmt.Println(receive2)
		//default:
		//	time.Sleep(10 * time.Second)
		//	fmt.Println("No channel is ready")
	}

}
