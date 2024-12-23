package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	wow := "W" + strings.Repeat("o", n) + "w!"

	fmt.Println(wow)
}
