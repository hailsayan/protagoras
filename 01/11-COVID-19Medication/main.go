package main

import "fmt"

func main() {
	var n, k, p, q int

	fmt.Scan(&n, &k, &p, &q)

	Shekarestan := n - k
	Namakestan := p - q

	if Shekarestan > Namakestan {
		fmt.Print("Shekarestan")
	} else if Shekarestan < Namakestan {
		fmt.Print("Namakestan")
	} else {
		fmt.Print("Equal")
	}
}
