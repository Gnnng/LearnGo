package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	v := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("The string \"%s\" has %d characters and its length is %d\n", v, utf8.RuneCountInString(v), len(v))
}
