package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func isTriangle(a, b, c int) string {
	if a > 0 && b > 0 && c > 0 && a+b+c == 180 {
		return "Yes"
	} else {
		return "No"
	}
}

func main() {
	defer writer.Flush()

	var a, b, c int
	scanf("%d %d %d\n", &a, &b, &c)
	printf("%s\n", isTriangle(a, b, c))
}
