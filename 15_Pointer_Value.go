package main

import "fmt"

type Fruit struct {
	price float64
}

// read and write
func (this *Fruit) RisePriceWithPointer() {
	this.price++
}

// read only, can not modify as we wish
func (this Fruit) RisePriceWithValue() {
	this.price++
}

// read only, suitable for getter
func (this Fruit) Price() float64 {
	return this.price
}

func main() {
	apple := Fruit{10}
	pear := Fruit{20}
	fmt.Println(apple.Price())
	fmt.Println(pear.Price())

	apple.RisePriceWithValue()
	fmt.Println(apple.Price())

	// no need to (&apple).RisePriceWithPointer
	apple.RisePriceWithPointer()
	fmt.Println(apple.Price())
}
