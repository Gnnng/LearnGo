package main

import "fmt"

type stack []int

func (s *stack) push(a int) {
	*s = append(*s, a)
}

func (s *stack) pop() (top int) {
	top = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

func (s *stack) String() (ret string) {
	return fmt.Sprint([]int(*s))
}

func main() {
	var aStack stack
	aStack.push(1)
	aStack.push(2)
	aStack.push(42)
	fmt.Println(&aStack) // need & to use our custom stringify method
	fmt.Println(aStack.pop())
	fmt.Println(&aStack)
}
