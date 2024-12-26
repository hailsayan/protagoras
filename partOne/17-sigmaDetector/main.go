package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriterSize(os.Stdout, math.MaxInt32)

func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func main() {
	var n, m int
	scanf("%d\n %d\n", &n, &m)

	var sum int
	for i := -10; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sum+=int(math.Pow(float64(i+j),3)/float64(j * j))
		}
	}
	printf("%d\n", sum)
	writer.Flush()
}
