package main

import "fmt"

func f() (ret int) {
	defer func() {
		ret++ // modified
		fmt.Println("Deferred")
	}()
	return func() int { fmt.Println("Returned"); return 0 }()
	//  will return 1
}

func un(s string) {
	fmt.Println("Leaving func", s)
}

func trace(s string) (ret string) {
	fmt.Println("Entering func", s)
	return s
}

func hello() {
	// The parameter of deferred function will be executed right away
	defer un(trace("hello"))
	fmt.Println("Hello")
}

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i) // Multiple deferred lines obey the rule of FIFO
		// will output 4 3 2 1 0
	}
	fmt.Println(f())
	hello()
}
