package main

import (
	"./stack"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Num int

type Operand struct {
	c     rune
	prior rune
	cal   func(n1, n2 Num) Num
}

var OpSet = map[rune]Operand{
	'+': Operand{'+', 0, func(n1, n2 Num) Num { return n1 + n2 }},
	'-': Operand{'-', 0, func(n1, n2 Num) Num { return n1 - n2 }},
	'*': Operand{'*', 1, func(n1, n2 Num) Num { return n1 * n2 }},
	'/': Operand{'/', 1, func(n1, n2 Num) Num { return n1 / n2 }},
	'%': Operand{'%', 1, func(n1, n2 Num) Num { return n1 % n2 }},
}

func isSameType(r1, r2 rune) bool {
	_, isOp1 := OpSet[r1]
	_, isOp2 := OpSet[r2]

	switch {
	case isOp1 && isOp2:
		fallthrough
	case r1 <= '9' && r1 >= '0' && r2 <= '9' && r2 >= '0':
		return true
	}
	return false
}

// MathSplitFunc is a split function for a Scanner that returns each token in
// a mathematic expression. This function only apply on pure-ascii string.
// Only the validation of character is checked. So return value may be "++++"
func MathSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces
	skipSpaces := 0
	for skipSpaces < len(data) && data[skipSpaces] == ' ' {
		skipSpaces++
	}

	// Scan until a different token or space
	for i := skipSpaces; i < len(data); i++ {
		if !isSameType(rune(data[skipSpaces]), rune(data[i])) {
			return i, data[skipSpaces:i], nil
		}
	}

	if atEOF && len(data) > skipSpaces {
		return len(data), data[skipSpaces:], nil
	}

	return skipSpaces, nil, nil
}

func main() {
	bio := bufio.NewReader(os.Stdin)
	s, _ := bio.ReadString('\n')

	//test_s :=  "10 + 20 + 3 + 4 - 10"
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(MathSplitFunc)

	numStack := stack.New()
	opStack := stack.New()

	for scanner.Scan() {
		token := scanner.Text()
		//fmt.Printf("Token: __%s__\n", token)
		if len(token) == 0 {
			break
		}
		if op, isOp := OpSet[rune(token[0])]; isOp { // Deal with operand
			// FIXME: Should not simply panic
			if len(token) > 1 {
				panic("not a valid operand")
			}
			for len(opStack) != 0 && OpSet[op.c].prior <= OpSet[rune(opStack[len(opStack)-1])].prior {
				n1 := numStack.Pop()
				n2 := numStack.Pop()
				calOp := opStack.Pop()
				numStack.Push(int(OpSet[rune(calOp)].cal(Num(n2), Num(n1))))
			}
			opStack.Push(int(op.c))
		} else { // Deal will number
			n, _ := strconv.Atoi(token)
			//fmt.Println("Convert to", n)
			numStack.Push(n)
		}
		//fmt.Println(numStack, opStack)
	}

	for len(opStack) != 0 {
		n1 := numStack.Pop()
		n2 := numStack.Pop()
		calOp := opStack.Pop()
		numStack.Push(int(OpSet[rune(calOp)].cal(Num(n2), Num(n1))))
	}

	fmt.Println(numStack.Pop())
}
