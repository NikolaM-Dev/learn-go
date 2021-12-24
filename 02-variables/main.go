package main

import "fmt"

func main() {
	var age int // variable declaration
	fmt.Println("My age is", age)

	age = 29 // assignment
	fmt.Println("My age is", age)

	age = 54 // assignment
	fmt.Println("My age is", age)

	var age2 int = 29 // variable declaration with initial value
	fmt.Println("My age is", age2)

	var age3 = 29 // type will be inferred
	fmt.Println("My age is", age3)

	var width, height int = 100, 50 // declaring multiple variables
	fmt.Println("width is", width, "height is", height)

	var width2, height2 = 100, 50 // "int" is dropped
	fmt.Println("width is", width2, "height is", height2)

	var (
		name    = "Nikola"
		age4    = 21
		height3 int
	)
	fmt.Println("My name is", name, ", age is", age4, "and height is", height3)

	count := 10 // short hand declaration
	fmt.Println("Count =", count)

	name2, age5 := "Nikola", 29 // short hand declaration
	fmt.Println("my name is", name2, "age is", age5)

	a, b := 20, 30 // declare variables a and b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b is already declared but c is new
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values to already declared variables b and c
	fmt.Println("changed b is", b, "c is", c)

	// d, e := 20, 30 // d and e declared
	// fmt.Println("a is", a, "b is", b)
	// d, e := 40, 50 // error, no new variables

	// age6 := 29      // age is int
	// age6 = "naveen" // error since we are trying to assign a string to a variable of type int
}
