package main

import "fmt"

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}

func printArray(h [3][2]string) {
	for _, v1 := range h {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	var a [3]int // int array with length 3
	a[0] = 12    // array index starts at 0
	a[1] = 78
	a[2] = 50
	fmt.Println(a)

	b := [3]int{12, 78, 70} // short hand declaration to create array
	fmt.Println(b)

	c := [3]int{12}
	fmt.Println(c)

	d := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(len(d))

	// e := [3]int{5, 78, 8}
	// var f [5]int
	// e = f // not possible since [3]int and [5]int are distinct types

	// Arrays are value types
	// Arrays in Go are value types and not reference types. This means
	// that when they are assigned to a new variable, a copy of the
	// original array is assigned to the new variable. If changes are
	// made to the new variable, it will not be reflected in the original
	//  array.
	e := [...]string{"USA", "China", "India", "Germany", "France"}
	f := e // a copy of a assigned to b
	f[0] = "Singapore"
	fmt.Println("e is ", e)
	fmt.Println("f is ", f)

	// Similarly when arrays are passed to functions as parameters,
	// they are passed by value and the original array is unchanged.
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num)
	fmt.Println("after passing to function ", num)

	// Length of an array
	g := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of g is", len(g))

	// Iterating arrays using range
	for i := 0; i < len(g); i++ { // looping from 0 to the length of the array
		fmt.Printf("%d the element of g is %.2f\n", i, g[i])
	}

	sum := float64(0)
	for i, v := range g { // range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of g", sum)

	// for _, v := range g { // ignores index
	// }

	// Multidimensional arrays
	h := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, // this comma is necessary. The compiler will complain if you omit this comma
	}
	printArray(h)

	var i [3][2]string
	i[0][0] = "apple"
	i[0][1] = "samsung"
	i[1][0] = "microsoft"
	i[1][1] = "google"
	i[2][0] = "AT&T"
	i[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printArray(i)

}
