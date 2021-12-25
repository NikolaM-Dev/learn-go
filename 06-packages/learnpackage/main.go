package main

import (
	"fmt"
	"log"

	"learnpackage/simpleinterest"
)

var p, r, t = -5000.0, 10.0, 1.0

func init() {
	println("Main package initializd")
	if p < 0 {
		log.Fatal("Principal is less than zero")
	}
	if r < 0 {
		log.Fatal("Rate of interest is less than zero")
	}
	if t < 0 {
		log.Fatal("Duration is less than zero")
	}
}

// The order of initialisation of the is as follows,

// 1. The imported packages are initialized first. Hence simpleinterest
//    package is initialized first and it's init method is called.
// 2. Package level variables p, r and t are initialized next.
// 3. init function is called in main.
// 4. main function is called at last.

func main() {
	fmt.Println("Simple interest calculation")

	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}
