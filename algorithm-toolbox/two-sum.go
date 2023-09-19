package main

import "fmt"

func CalculateTwoSum(a int, b int) int {
	return a + b
}

func main() {
	//get input
	var a, b int
	fmt.Scan(&a, &b)
	var result = CalculateTwoSum(a, b)
	fmt.Println(result)
}
