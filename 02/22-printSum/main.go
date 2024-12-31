package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// n * (n + 1) / 2

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func main() {
	defer writer.Flush()
	var n int
	scanf("%d\n", &n)

	result := ""
	for i := 1; i <= n; i++ {
		result += strconv.Itoa(i)
		if i < n {
			result += " + "
		}
	}

	sum := n * (n + 1) / 2
	result += " = " + strconv.Itoa(sum)

	printf("%s\n", result)
}

