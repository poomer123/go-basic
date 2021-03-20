package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func next(num int) (int, int) {
	return num + 1, num + 2
}

func main() {

	result := add(10, 20)

	x, y := next(10)

	fmt.Println(result, x, y)
}
