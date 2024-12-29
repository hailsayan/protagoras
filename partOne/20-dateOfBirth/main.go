package main

import "fmt"

func dateOfBirth(input string) string {
	year := input[:2]
	date := input[2:]

	return "saal:" + year + "\nmaah:" + date
}

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Println(dateOfBirth(input))
}
