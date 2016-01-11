package stack

import "fmt"

type Stack []int

func New() Stack {
    return make(Stack, 0)
}

func (s *Stack) Push(a int) {
	*s = append(*s, a)
}

func (s *Stack) Pop() (top int) {
	top = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

func (s *Stack) String() (ret string) {
	return fmt.Sprint([]int(*s))
}

