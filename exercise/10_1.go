package main

import "fmt"

func fibonacci(n int) (fib []int) {
	if n <= 0 {
		return
	}
	fib = append(fib, 1)
	if n == 1 {
		return
	}
	fib = append(fib, 1)
	if n == 2 {
		return
	}
	for x1, x2, i := 1, 1, 3; i <= n; i++ {
		xn := x1 + x2
		x1, x2 = x2, xn
		fib = append(fib, xn)
	}
	return
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}
}
