package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(s)
	for _, v := range s {
		t = string(v) + t
	}
	fmt.Println(t)
}
