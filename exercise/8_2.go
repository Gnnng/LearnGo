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
	for i, v := range *s {
		ret = ret + " " + fmt.Sprintf("[%d:%d]", i, v)
	}
	return ret[1:]
}

func main() {
	var aStack stack
	aStack.push(1)
	aStack.push(2)
	aStack.push(42)
	fmt.Println(&aStack)
	fmt.Println(aStack.pop())
	fmt.Println(&aStack)
}
