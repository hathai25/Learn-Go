package main

import (
	"fmt"
	"sync"
)

var on sync.Once
var wg sync.WaitGroup

func setup() {
	//init appears only once
	fmt.Println("Init")
}

func doStuff() {
	on.Do(setup)
	fmt.Println("Hello")
	wg.Done()
}

func main() {
	wg.Add(2)
	go doStuff()
	go doStuff()
	wg.Wait()
}
