package main

import (
	"fmt"
	"strconv"
)

/*
FizzBuzz

In this problem, you should display a list from 1 to 100, one on each line, with the following exceptions:
Numbers divisible by 3 should appear as 'Fizz' instead of number;
Numbers divisible by 5 should appear as 'Buzz' instead of number;
Numbers divisible by 3 and 5 should appear as' FizzBuzz 'instead of number'.
*/

func fizzBuzz(i int) string {
	switch {
	case i%5 == 0 && i%3 == 0:
		return "FizzBuzz"
	case i%3 == 0:
		return "Fizz"
	case i%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(i)
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzBuzz(i))
	}
}
