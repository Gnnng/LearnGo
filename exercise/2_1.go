package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		numFlag := true
		if i%3 == 0 {
			fmt.Printf("Fizz")
			numFlag = false
		}
		if i%5 == 0 {
			fmt.Printf("Buzz")
			numFlag = false
		}
		if numFlag {
			fmt.Print(i)
		}
		if i%10 == 0 {
			fmt.Print("\n")
		} else {
			fmt.Println()
		}
	}
}
