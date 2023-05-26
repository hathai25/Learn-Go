package main

import "fmt"

type student struct {
	id, age int
	name    string
	major   string
}

func increment(num *int) {
	*num++
}

func main() {
	var ptr *int
	var my_int = 10
	fmt.Println(my_int)
	ptr = &my_int
	increment(ptr)
	fmt.Println(my_int)

	var Thai = student{1, 20, "Thai", "Computer Science"}
	var ptrThai *student
	var secPtrThai **student
	ptrThai = &Thai
	secPtrThai = &ptrThai
	fmt.Println(ptrThai)
	ptrThai.major = "Software Engineering"
	fmt.Println(ptrThai)
	(*secPtrThai).major = "Computer Engineering"
	fmt.Println(ptrThai)
}
