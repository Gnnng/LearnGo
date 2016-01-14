package main

import "fmt"

func main() {
	// use new to apply allocation of memory and use zero-value of its type to
	// initialize it, the return value is a pointer
	var p *[]int = new([]int)
	*p = make([]int, 100)

	// make simpily returns the value ofter initialization
	// make must be applied on array, slice and map
	v := make([]int, 100)

	fmt.Println()
}
