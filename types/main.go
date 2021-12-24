package main

import (
	"fmt"
	"unsafe"
)

/*
The following are the basic types available in Go

- bool
- Numeric Types
    - int8, int16, int32, int64, int
    - uint8, uint16, uint32, uint64, uint
    - float32, float64
    - complex64, complex128
    - byte
    - rune
- string
*/

func main() {
	a := true
	b := false

	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)

	d := a || b
	fmt.Println("d:", d)

	/*
		Signed integers

		int8: represents 8 bit signed integers
		size: 8 bits
		range: -128 to 127

		int16: represents 16 bit signed integers
		size: 16 bits
		range: -32768 to 32767

		int32: represents 32 bit signed integers
		size: 32 bits
		range: -2147483648 to 2147483647

		int64: represents 64 bit signed integers
		size: 64 bits
		range: -9223372036854775808 to 9223372036854775807

		int: represents 32 or 64 bit integers depending on the underlying platform. You should generally be using int to represent integers unless there is a need to use a specific sized integer.
		size: 32 bits in 32 bit systems and 64 bit in 64 bit systems.
		range: -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems
	*/
	var a2 int = 89
	b2 := 95
	fmt.Println("value of a is", a2, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a2, unsafe.Sizeof(a2))   // type and size of a2
	fmt.Printf("\ntype of b is %T, size of b is %d", b2, unsafe.Sizeof(b2)) // type and size of b2

	/*
		Unsigned integers
		uint8: represents 8 bit unsigned integers
		size: 8 bits
		range: 0 to 255

		uint16: represents 16 bit unsigned integers
		size: 16 bits
		range: 0 to 65535

		uint32: represents 32 bit unsigned integers
		size: 32 bits
		range: 0 to 4294967295

		uint64: represents 64 bit unsigned integers
		size: 64 bits
		range: 0 to 18446744073709551615

		uint : represents 32 or 64 bit unsigned integers depending on the underlying platform.
		size : 32 bits in 32 bit systems and 64 bits in 64 bit systems.
		range : 0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems
	*/

	/*
		Floating point types

		float32: 32 bit floating point numbers
		float64: 64 bit floating point numbers
	*/
	a3, b3 := 5.67, 8.97
	fmt.Printf("type of a3 %T b3 %T\n", a3, b3)
	sum := a3 + b3
	diff := a3 - b3
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)

	/*
		Complex types
		complex64: complex numbers which have float32 real and imaginary parts
		complex128: complex numbers with float64 real and imaginary parts
	*/
	e1 := complex(5, 7) // 5 real, 7 imaginary
	e2 := 8 + 27i       // 8 real, 27 imaginary
	eadd := e1 + e2
	fmt.Println("sum:", eadd)

	emul := e1 * e2
	fmt.Println("product:", emul)

	/*
		Other numeric types
		byte is an alias of uint8
		rune is an alias of int32
	*/

	/*
		string type
		Strings are a collection of bytes in Go
	*/
	first := "Nikola"
	last := "Merchan"
	name := first + " " + last
	fmt.Println("My name is", name)

	/*
		Type Conversion
		Go is very strict about explicit typing. There is no automatic type promotion or conversion. Let's look at what this means with an example.
	*/
	i := 55   // int
	j := 67.8 // float64
	// sum := i + j // int + float64 not allowed
	sum2 := i + int(j) // j is convert to int
	fmt.Println(sum2)

	k := 10
	var h float64 = float64(k) //this statement will not work without explicit conversion
	fmt.Println("h", h)
}
