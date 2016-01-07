package main

import (
	"fmt"
)

func main() {
	v := "asSASA ddd dsjkdsjs dk"
	d := v[:]
	s := "abc" + v[3:]
	fmt.Println(s)
}
