package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	left, right *sync.Mutex
}

func Eat(c chan int, phi Philosopher, num int, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		c <- 0
		fmt.Println("starting to eat", num)
		time.Sleep(time.Millisecond * 500)
		phi.left.Lock()
		phi.right.Lock()

		phi.left.Unlock()
		phi.right.Unlock()
		fmt.Println("finishing eating", num)
		time.Sleep(time.Millisecond * 500)
		<-c
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int, 2)
	var chopsticks []sync.Mutex
	for i := 0; i < 5; i++ {
		chopsticks = append(chopsticks, sync.Mutex{})
	}

	var philosophers []Philosopher
	for i := 0; i < 5; i++ {
		philosophers = append(philosophers, Philosopher{left: &chopsticks[i], right: &(chopsticks[(i+1)%5])})
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Eat(c, philosophers[i], i+1, &wg)
	}

	wg.Wait()
}
