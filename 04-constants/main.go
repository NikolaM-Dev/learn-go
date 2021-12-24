package main

import (
	"fmt"
)

func main() {
	const a = 50
	fmt.Println(a)

	// Declaring a group of constants
	const (
		name    = "Nikola"
		age     = 21
		country = "Colombia"
	)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)

	// The value of a constant should be known at compile time.
	// var c = math.Sqrt(4) // allowed
	// const d = math.Sqrt(4) // not allowed

	const trueConst = true
	type myBool bool
	// var defaultBool = trueConst       // allowed
	// var customBool myBool = trueConst // allowed
	// defaultBool = customBool          // not allowed
}
