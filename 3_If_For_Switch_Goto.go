package main

import "fmt"

func main() {
	var x int

	fmt.Scanf("%d", &x)

	// Section 1: If Statement
	// What "if" looks like
	if x > 0 {
		fmt.Println("x > 0")
	} else if x < 0 {
		fmt.Println("x < 0")
	} else {
		fmt.Println("x = 0")
	}

	// or
	if fmt.Scanf("%d", &x); x > 0 {
		fmt.Println("x > 0")
	} else if x < 0 {
		fmt.Println("x < 0")
	} else {
		fmt.Println("x = 0")
	}

	// Section 2: For Statement
	// for loop
	for i := 0; i < 3; i++ {
		fmt.Println("Current i = ", i)
	}

	// or without initialization and continuation
	// looks more like a while statement
	i := 0
	for i < 3 {
		fmt.Println("Current i = ", i)
		i++
	}

	// and dead loop
	/*
	   for {}

	*/

	// for with break or continue is same as its actions in C
	// or
outerLoop:
	for i := 0; i < 3; i++ { // i++ and i-- are not expressions but statements
	innerLoop:
		for j := 0; j < 5; j++ {
			if j > 4 {
				break innerLoop // by default, or write as "continue outerLoop"
			}
			if j > 2 {
				break outerLoop
			}
		}
	}

	// keyword range
	for pos, c := range "Hello, 中国" {
		fmt.Printf("%c in position %v\n", c, pos)
	}

	// Section 3: Switch Statement
	var c int
	fmt.Scanf("%d", &c)
	switch c {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fallthrough
	case 4:
		fmt.Println("3 & 4")
	case 5, 6, 7:
		fmt.Println("5, 6 & 7")
	}

	switch k := 10; {
	case k < 5:
		fmt.Println("k < 5")
	case k > 5:
		fmt.Println("k > 5")
	}

	// goto
	accessed := false
aLabel:
	if accessed {
		fmt.Println("End")
		return
	}
	fmt.Println("Not end")
	accessed = true
	goto aLabel
}
