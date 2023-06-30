package main

import "sync"

var wg sync.WaitGroup

func doStuff(ch1 chan int, ch2 chan int) {
	defer wg.Done()

	<-ch1
	ch2 <- 1
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	wg.Add(2)
	go doStuff(ch1, ch2)
	go doStuff(ch2, ch1)
	wg.Wait()
}
