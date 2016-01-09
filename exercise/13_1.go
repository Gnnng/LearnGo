package main

import "fmt"

func bubbleSort(slice []int) {
	for swapped := true; swapped; {
		swapped = false
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] < slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				swapped = true
			}
		}
	}
}

func main() {
	a := []int{1, 0, 9, -1, -3, -10, 12}
	bubbleSort(a)
	fmt.Println(a)
}
