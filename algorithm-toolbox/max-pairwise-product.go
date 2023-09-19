package main

import (
	"fmt"
	"slices"
)

func main() {
	var n int
	fmt.Scan(&n)
	var numbers []int
	for i := 0; i < n; i++ {
		var number int
		fmt.Scan(&number)
		numbers = append(numbers, number)
	}

	var max1 = slices.Max(numbers)

	numbers = slices.DeleteFunc(numbers, func(i int) bool {
		return i == max1
	})
	var max2 = slices.Max(numbers)
	fmt.Println(max1 * max2)

}
