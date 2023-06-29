package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

func Sort(id int, numbers []int, wg *sync.WaitGroup) {
	defer wg.Done()

	if len(numbers) > 0 {
		fmt.Printf("Goroutine [%d] will sort %v\n", id, numbers)
		sort.Ints(numbers)
	} else {
		fmt.Printf("Goroutine [%d] has nothing to sort\n", id)
	}
}

func main() {
	numbers := make([]int, 0)

	for {
		var valStr string

		fmt.Print("Enter an integer to sort or 'x' when done: ")
		fmt.Scanf("%s\n", &valStr)

		if valStr == "x" {
			break
		}

		val, err := strconv.Atoi(valStr)
		if err == nil {
			numbers = append(numbers, val)
			fmt.Printf("Added number [%d]\n", val)
		} else {
			fmt.Printf("Invalid number [%s]\n", valStr)
		}
	}

	var wg sync.WaitGroup
	for routineID := 0; routineID < 4; routineID++ {
		// split slice in 4
		var startIdx, endIdx int

		startIdx = (len(numbers) * routineID) / 4
		if routineID < 3 {
			endIdx = (len(numbers) * (routineID + 1)) / 4
		} else {
			endIdx = len(numbers)
		}

		wg.Add(1)
		go Sort(routineID, numbers[startIdx:endIdx], &wg)
	}

	wg.Wait()
	fmt.Printf("All sort goroutines ended. Main will now sort %v\n", numbers)

	// sort the full slice now
	sort.Ints(numbers)

	fmt.Println(numbers)
}
