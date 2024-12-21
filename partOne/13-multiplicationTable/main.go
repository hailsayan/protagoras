package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%d ", i*j)
		}
		fmt.Println()
	}
}
