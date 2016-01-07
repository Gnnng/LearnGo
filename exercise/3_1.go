package main

import "fmt"

func main() {
	for count, i := 0, 1; count < 100; i++ {
		count += i
		for j := 0; j < i; j++ {
			fmt.Print("A")
		}
		fmt.Println()
	}
}
