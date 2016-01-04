package main // semicolon is optional

import "fmt"

// Declare null as default value
var (
	a int    // = 0
	b string // = ""
	c bool   // = false
	d error  // = nil
)

// More number types
var (
	aint        int   // and uint
	aint8       int8  // and uint8/byte
	aint16      int16 // and uint16
	aint32      int32 // and uint32
	arune       rune  // alias of int32, for special usage
	aint64      int64 // and uint64
	afloat32    float32
	afloat64    float64
	acomplex64  complex64 // = 0 + 0i
	acomplex128 complex128
)

// const can be number, string or bool
const (
	aa = 42
	bb = "foobar"
	cc = true
	// or other constant literal value
	dd = 077  // octal
	ee = 0xEF // hex
	ff = 1e9
	gg = 314e-2
)

// const can be initialized by iota enumerator
const (
	_   = iota
	KiB = 1 << (10 * iota)
	MiB // implicitly repeated
	GiB // implicitly repeated
	// Skip TiB
	PiB = 1 << (10 + 10*iota)
	EiB // implicitly repeated, will be same as PiB
)

// Different Variable Declaration Method

// Method 1: Default Initialization
var x, y, z int

// Method 2: With Initializer (type omitted)
var xx, yy, zz = 10, true, "Good"

func main() {
	// Method 3: Use "short assignment" (:=) statement
	xxx, yyy, zzz := 10, true, "Good" // use only in func

	fmt.Println(xxx)
	fmt.Println(yyy)
	fmt.Println(zzz)

	// Multi-line string, check out their difference
	var s = "This is an" +
		" example of multi-line string"

	var ss = `This is another example,
        which uses antiquote`

	fmt.Println(s)
	fmt.Println(ss)
	fmt.Println("End")
}
