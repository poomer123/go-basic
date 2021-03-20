package main

import "fmt"

func main() {
	const pi = 3.14
	// can't do this pi = 10

	const (
		one = iota + 1
		two
		three
		_
		_
		six
	)
	fmt.Println(one, two, three, six)
}
