package main

import "fmt"

func plusXFactory(x int) (f func(int) int) {
	return func(a int) int { return x + a }
}

func main() {
	plusTwo := plusXFactory(2)
	plusTen := plusXFactory(10)
	fmt.Println(plusTwo(3), plusTen(2))
}
