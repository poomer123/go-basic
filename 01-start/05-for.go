package main

import "fmt"

func main() {

	// example 1
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}

	// example 2
	n := 1
	for n <= 10 {
		fmt.Println(n)
		n++
	}
}
