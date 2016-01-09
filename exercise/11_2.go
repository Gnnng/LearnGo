package main

import "fmt"

func slice_map(f func(int) int, slice []int) (ret []int) {
	ret = make([]int, len(slice))
	for i, v := range slice {
		ret[i] = f(v)
	}
	return
}

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice_map(func(x int) int { return 2 * x }, a[:]))
}
