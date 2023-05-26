package main

import "fmt"

func closure() func() string {
	names := []string{"sad", "boi", "hours"}
	index := -1
	return func() string {
		index++
		return names[index]
	}

}

func main() {
	f := closure()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
