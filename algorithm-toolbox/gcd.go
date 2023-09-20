package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var remainder = a % b
	var gcd = 1
	for remainder != 0 {
		a = b
		b = remainder
		remainder = a % b
		gcd = b
	}
	fmt.Println(gcd)
}
