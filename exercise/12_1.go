package main

import "fmt"

func max(slice []int) (ret int) {
	ret = -int(^uint(0)>>1) - 1 // minimal of int (wether 32-bit or 64 bit)
	for _, v := range slice {
		if v > ret {
			ret = v
		}
	}
	return
}

func main() {
	fmt.Println(max([]int{-2, -3, 1, 2, 3}))
}
