package main

import (
	"./stack"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func opLessThan(op1, op2 int) bool {
	opLevel := map[rune]int{
		'+': 0, '-': 0, '*': 1, '/': 1,
	}

	return opLevel[rune(op1)] <= opLevel[rune(op2)]
}

func cal(op, num1, num2 int) (result int) {
	switch rune(op) {
	case '+':
		return num1 + num2
	case '-':
		return num1 - num2
	case '*':
		return num1 * num2
	case '/':
		return num1 / num2
	}
	return 0
}

func main() {
	bio := bufio.NewReader(os.Stdin)
	s, _ := bio.ReadString('\n')

	numStack := stack.New()
	opStack := stack.New()

	last := 0
	isOp := map[rune]bool{'+': true, '-': true, '*': true, '/': true}
	for i, c := range s {
		if isOp[c] {
			num, err := strconv.Atoi(strings.TrimSpace(s[last:i]))
			if err != nil {
				panic(err)
			}
			numStack.Push(num)
			for len(opStack) != 0 {
				if opLessThan(int(c), opStack[len(opStack)-1]) {
					num2 := numStack.Pop()
					num1 := numStack.Pop()
					numStack.Push(cal(opStack.Pop(), num1, num2))
				} else {
					break
				}
			}
			opStack.Push(int(c))
			last = i + 1
		}
	}
	num, err := strconv.Atoi(strings.TrimSpace(s[last:len(s)]))
	if err != nil {
		panic(err)
	}
	numStack.Push(num)
	for len(opStack) != 0 {
		num2 := numStack.Pop()
		num1 := numStack.Pop()
		numStack.Push(cal(opStack.Pop(), num1, num2))
	}
	fmt.Println(numStack.Pop())
}
