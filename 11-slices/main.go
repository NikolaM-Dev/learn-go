package main

import "fmt"

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	countriesNeeded := countries[:len(countries)-2]
	countriesCpy := make([]string, len(countriesNeeded))
	copy(countriesCpy, countriesNeeded) // copies neededCountries to countriesCpy
	return countriesCpy
}

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] // create a slice from a[1] to a[3]
	fmt.Println(b)

	c := []int{6, 7, 8} // creates and array and returns a slice reference
	fmt.Println(c)

	// Modifying a slice
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)

	// When a number of slices share the same underlying array,
	// the changes that each one makes will be reflected in the array.
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] // creates a slice wich contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before chage 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after moditication to slice nums2", numa)

	// Length and capacity of a slice
	fruitarray := [...]string{
		"apple",
		"orange",
		"grape",
		"mango",
		"water melon",
		"pine apple",
		"chikoo",
	}

	fruitslice := fruitarray[1:3]
	// length of fruitslice is 2 and capacity is 6
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))

	// re-slicing fruitslice till its capacity
	fruitslice = fruitslice[:cap(fruitslice)]
	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity", cap(fruitslice))

	// Creating a slice using make
	i := make([]int, 5, 5)
	fmt.Println(i)

	// Appending to a slice
	cars := []string{"Ferrari", "Honda", "Ford"}
	// capacity of cars is 3
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyota")
	// capacity of cars is doubled to 6
	fmt.Println("cars:", cars, "has new length", len(cars), "and capaicty", cap(cars))

	var names []string // zero value of a slice is nil
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}

	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)

	// Passing a slice to a function
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)
	// function modifies the slice
	// modifications are visible outside
	fmt.Println("slice after function call", nos)

	// Multidimensional slices
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}

	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}

	// Memory Optimisation
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}
