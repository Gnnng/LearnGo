package main

import "fmt"

// both R and S implement the interface I
type S struct{ i int }

func (p *S) Get() int  { return p.i }
func (p *S) Put(v int) { p.i = v }

type R struct{ i int }

func (p *R) Get() int  { return p.i }
func (p *R) Put(v int) { p.i = v }

type I interface {
	Get() int
	Put(int)
}

type J interface {
	I
}

type Anything interface{}

func f(p I) {
	// Type switch, .(type) can only be used in type switch
	switch p.(type) {
	case *S:
		fmt.Println("S")
	case *R:
		fmt.Println("R")
	}
}

func g(x Anything) {
	if t, ok := x.(I); ok { //if x implements interface I, ok = true
		fmt.Println(t, "implements I")
	}
}

func main() {
	f(new(S))
	f(new(R))
	s := S{123}
	g(&s)
	g(J(&s))
}
