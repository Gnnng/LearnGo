package main

import "fmt"

func plusTwoFactory() (f func(int) int) {
	return func(x int) int { return x + 2 }
}

func main() {
	plusTwo := plusTwoFactory()
	fmt.Println(plusTwo(3))
}
