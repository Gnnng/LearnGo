package main

import "fmt"

func average(s []float64) (result float64) {
	for _, i := range s {
		result += i
	}
	return result / float64(len(s))
}

func main() {
	fmt.Println(average([]float64{1.1, 1.2, 1.3}))
}
