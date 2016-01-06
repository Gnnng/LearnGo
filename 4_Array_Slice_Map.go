package main

import "fmt"

func main() {
	var a [100]int
	var b [...]string

	cc := [...]int{1, 2, 3}

	fmt.Println(a, b, cc)
}
