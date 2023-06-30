package main

import "sync"

var i int = 0

var mut sync.Mutex
var wg sync.WaitGroup

func inc() {
	mut.Lock()
	i = i + 1
	mut.Unlock()
	wg.Done()
}

func main() {
	wg.Add(2)
	go inc()
	go inc()
	wg.Wait()
	println(i)
}
