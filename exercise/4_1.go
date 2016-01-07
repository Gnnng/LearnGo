package main

import "fmt"

func average(s ...float64) (result float64) {
	for _, i := range s {
		result += i
	}
	return result / float64(len(s))
}

func main() {
	fmt.Println(average(1.1, 1.2, 1.3))
	fmt.Println(average(1.1, 1.2, 1.3, 1.4, 1.5))
	fmt.Println(average(1.1, 1.2, 1.3, 10, 20))
}
