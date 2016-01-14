package main

import "fmt"

type Namecard struct {
	Name string // public
	Age  uint   // public

	id uint // private
}

func (n *Namecard) String() (ret string) {
	ret += fmt.Sprintln("\n==============================")
	ret += fmt.Sprintln("\tName:", n.Name)
	ret += fmt.Sprintln("\tAge:", n.Age)
	ret += fmt.Sprintln("==============================")
	return ret
}

// String() method of Namecard will be derived by Anonymous
type Anonymous struct {
	// anonymous field
	string
	*Namecard
}

func main() {
	x := new(Namecard)
	x.Name = "John"
	x.Age = 18
	x.id = 1

	y := Namecard{"Tom", 19, 2}

	z := Namecard{id: 3, Name: "Bob", Age: 20}

	a := Anonymous{"non", x}

	fmt.Printf("%v, %+v, %#v\n", x, x, x)
	fmt.Printf("%v, %+v, %#v\n", y, y, y)
	fmt.Printf("%v, %+v, %#v\n", z, z, z)
	// attention: `func (n *Namecard) String()` will be called in "%v" and "%+v"
	fmt.Printf("%v, %+v, %#v\n", a, a, a)
}
