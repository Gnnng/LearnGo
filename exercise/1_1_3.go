package main

import "fmt"

func main() {
	numbers := [...]string{"zero", "one", "two", "three"}
	i := 0
loop:
	fmt.Println(numbers[i])
	if i++; i < len(numbers) {
		goto loop
	}
}
