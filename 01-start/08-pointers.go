package main

import "fmt"

func main() {
	i := 20
	p := &i

	fmt.Println(p, *p)
}
