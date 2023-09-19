package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var numbers []int
	for i := 0; i < n; i++ {
		var number int
		fmt.Scan(&number)
		numbers = append(numbers, number)
	}

	var max1 int = 0
	var max2 int = 0
	for _, number := range numbers {
		if number > max1 {
			max2 = max1
			max1 = number
		} else if number > max2 {
			max2 = number
		}
	}
	fmt.Println(max1 * max2)

}
