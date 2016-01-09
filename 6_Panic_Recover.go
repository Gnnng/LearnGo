package main

import "fmt"

func panicDetector(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil { // recover() is only valid in deferred function
			b = true
			fmt.Println(x)
		}
	}()
	f()
	return
}

func generatePanic() {
	panic("A panic")
}

func main() {
	if panicDetector(generatePanic) {
		fmt.Println("Panic Generated")
	}
}
