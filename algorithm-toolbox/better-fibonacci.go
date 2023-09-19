package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n == 1 || n == 2 {
		fmt.Println(1)
		return
	}
	if n > 2 {
		var numbers []int = make([]int, n+1)
		numbers[0] = 1
		numbers[1] = 1
		for i := 2; i < n+1; i++ {
			numbers[i] = numbers[i-1] + numbers[i-2]
		}
		fmt.Println(numbers[n])
	}
}
