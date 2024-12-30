package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	maxAmount := 0
	maxName := ""

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Split(line, " ")
		name := parts[0]
		amount, _ := strconv.Atoi(parts[1])

		if amount > maxAmount {
			maxAmount = amount
			maxName = name
		}
	}

	fmt.Println(maxName)
}
