package main

import "fmt"

func fibonacci(index int) int {
	if index == 1 || index == 2 {
		return 1
	}
	if index > 2 {
		return fibonacci(index-1) + fibonacci(index-2)
	}
	return 0
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(fibonacci(i))
	}
}
