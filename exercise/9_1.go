package main

import "fmt"

func printInt(s ...int) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	printInt(1, 2, 3, 4, 5)
	// or
	printInt([]int{1, 2, 3, 4, 5}...)
}
