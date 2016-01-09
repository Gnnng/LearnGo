package main

import "fmt"

func nOrder(a, b int) (aa, bb int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

func main() {
	fmt.Println(nOrder(7, 2))
	fmt.Println(nOrder(2, 7))
}
