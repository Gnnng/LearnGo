package main

import "fmt"

func main() {
	// create 1-dimension array with default value
	var a [5]int
	var b [5]string

	// short assignment
	aa := [5]int{1, 2, 3} // aa[3] ~ aa[9] are 0
	bb := [...]string{"good", "bad"}
	cc := [5]int{1, 2}

	fmt.Println(a, b)
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)

	// create 2-dimension array
	a2 := [2][2]int{}
	linesOfText := [][]byte{[]byte("line 1"), []byte("line 2")}
	//a2 := [2][2]int  // not valid

	fmt.Println("a2 is: ", a2)
	fmt.Println("linesOfText", linesOfText)

	// slice creation
	sl := make([]int, 5)
	//[ ][ ][ ][ ][ ]
	// | sl
	//             | len: 5
	//             | cap: 5
	fmt.Println("sl len:", len(sl), "cap:", cap(sl))
	sl = sl[:3]
	//[ ][ ][ ][ ][ ]
	// | sl
	//       | len: 3
	//             | cap: 5
	fmt.Println("sl len:", len(sl), "cap:", cap(sl))
	sl = sl[2:]
	//[ ][ ][ ][ ][ ]
	//       | sl
	//       | len: 1
	//             | cap: 3
	fmt.Println("sl len:", len(sl), "cap:", cap(sl))
	sl = append(sl, 10)
	fmt.Println("sl len:", len(sl), "cap:", cap(sl))

	sl = append(sl, []int{20, 30, 40}...) // must use three dots
	fmt.Println("sl len:", len(sl), "cap:", cap(sl))

	m := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4, // must contain last dot
	}

	fmt.Println(m["zero"])

	m["five"] = 5 // add new item
	fmt.Println(m)

	// non-existent key will return default value
	fmt.Println(m["six"])

	for i := 0; i < 3; i++ {
		switch i {
		case 0: // do nothing, six not in m
		case 1:
			m["six"] = 6 // add six
		case 2:
			delete(m, "six") //delete six
		}
		if value, ok := m["six"]; ok {
			fmt.Println("Have six, which is", value)
		} else {
			fmt.Println("Not have six")
		}
	}
}
