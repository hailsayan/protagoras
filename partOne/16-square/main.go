package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	x := 1
	for x <=n {
		x *= 2
	}
	fmt.Print(x)
}
