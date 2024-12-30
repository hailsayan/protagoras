package main

import "fmt"

func freedomOfSpeech(n int) string{
	if n%2 == 1{
		return "Payin Barare"
	} else {
		return "Bala Barare"
	}
}

func main() {
	var n int
	fmt.Scanf("%d ", &n)

	fmt.Println(freedomOfSpeech(n))
}