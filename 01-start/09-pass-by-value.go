package main

import "fmt"

func inc(num *int) {
	*num++
}

func main() {
	i := 20

	inc(&i)
	fmt.Println(i)
}
