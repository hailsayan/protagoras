package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	maxDigit := 0
	for i := 0; i < n; i++ {
		var digit int
		fmt.Scan(&digit)
		if maxDigit < digit {
			maxDigit = digit
		}
	}

	fmt.Println(maxDigit)
}
