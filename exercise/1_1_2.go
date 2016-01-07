package main

import "fmt"

func main() {
	i := 0
loop:
	fmt.Println(i)
	if i++; i < 10 {
		goto loop
	}
}
