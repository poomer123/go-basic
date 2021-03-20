package main

import "fmt"

func main() {

	day := "Saturday"

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Workday")
	}
}
