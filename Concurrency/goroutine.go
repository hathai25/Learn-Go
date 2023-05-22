package main

import (
	"fmt"
	"time"
)

/*
A Goroutine is a function or method which executes independently
and simultaneously in connection with any other Goroutines present in your program.

You can consider a Goroutine like a light weighted thread.
The cost of creating Goroutines is very small as compared to the thread.

Every program contains at least a single Goroutine and that Goroutine is known as the main Goroutine.
*/

func showConnection(message string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println(message, i)
	}
}

func main() {
	go showConnection("my routine")
	//when a new Goroutine executed, the Goroutine call return immediately
	fmt.Println("Next line")
	/* The control does not wait for Goroutine to complete their execution just like normal function they always
	move forward to the next line after the Goroutine call and ignores the value returned by the Goroutine. */

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			fmt.Println("Anonymous Goroutine")
		}
	}()

	showConnection("normal") // main Goroutine

}
