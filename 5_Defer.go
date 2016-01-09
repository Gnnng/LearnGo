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

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i) // Multiple deferred lines obey the rule of FIFO
		// will output 4 3 2 1 0
	}
	fmt.Println(f())

}
