package main

import "fmt"

func change(val *int) {
	*val = 55
}

func hello() *int {
	i := 5
	return &i
}

func modify(sls []int) {
	sls[0] = 90
}

func main() {
	// What is a pointer?
	// A pointer is a variable that stores the memory address of
	// another variable.

	// Declaring pointers
	b := 255
	// The & operator is used to get the address of a variable
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	// Zero value of a pointer
	c := 25
	var d *int
	if d == nil {
		fmt.Println("d is", d)
		d = &c
		fmt.Println("d after initialization is", d)
	}

	// Creating pointers using the new function
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)

	// Dereferencing a pointer
	e := 255
	f := &e
	fmt.Println("address of e is", f)
	fmt.Println("value of e is", *f)
	*f++
	fmt.Println("New value of e is ", e)

	// Passing pointer to a function
	g := 58
	fmt.Println("value of g before function call is", g)
	h := &g
	change(h)
	fmt.Println("value of g after function call is", g)

	// Returning pointer from a function
	i := hello()
	fmt.Println("Value of i", *i)

	// Do not pass a pointer to an array as an argument to a
	// function. Use slice instead.
	// j := [3]int{89, 90, 91}
	// modify(&j)
	// fmt.Println(j)

	// Although this way of passing a pointer to an array as an
	// argument to a function and making modification to it works,
	// it is not the idiomatic way of achieving this in Go.
	// We have slices for this.
	j := [3]int{89, 90, 91}
	modify(j[:])
	fmt.Println(j)

	// Go does not support pointer arithmetic
	// b2 := [...]int{109, 110, 111}
	// p := &b2
	// p++
}
