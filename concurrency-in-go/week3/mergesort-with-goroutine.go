package main

import (
	"fmt"
	"time"
)

func sortArr(arr []int) []int {
	//bubble sort the array
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr

}

// channel that receive an array of int and return a sorted array of int using sortArr
func sortArrChannel(arr []int, c chan []int) {
	fmt.Println("receive array: ", arr)
	c <- sortArr(arr)
}

func main() {
	start := time.Now()
	//make an array of int
	arr := []int{10, 9, 8, 7, 6, 1, 2, 3, 5, 4, 11}
	//create 4 sub array of arr
	arr1 := arr[:len(arr)/4]
	arr2 := arr[len(arr)/4 : len(arr)/2]
	arr3 := arr[len(arr)/2 : 3*len(arr)/4]
	arr4 := arr[3*len(arr)/4:]
	//create 4 channel of int
	c1 := make(chan []int)
	c2 := make(chan []int)
	c3 := make(chan []int)
	c4 := make(chan []int)
	//create 4 goroutine that sort the sub array
	go sortArrChannel(arr1, c1)
	go sortArrChannel(arr2, c2)
	go sortArrChannel(arr3, c3)
	go sortArrChannel(arr4, c4)
	//receive the sorted sub array
	sortedArr1 := <-c1
	sortedArr2 := <-c2
	sortedArr3 := <-c3
	sortedArr4 := <-c4
	//sort the merged sub array
	sortedArr := sortArr(append(append(sortedArr1, sortedArr2...), append(sortedArr3, sortedArr4...)...))
	fmt.Println(sortedArr)
	fmt.Println(time.Since(start))
}
