package main

import "fmt"

func calculateBill(price int, no int) int {
	var totalPrice = price * no
	return totalPrice
}

// If consecutive parameters are of the same type,
// we can avoid writing the type each time and it is
// enough to be written once at the end
func calculateBill2(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

// Multiple return values
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

// Named return values
// It is possible to return named values from a function. If a
// return value is named, it can be considered as being declared as
// a variable in the first line of the function.

func rectProps2(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // no explicit return value
}

// Blank Identifier
// _ is known as the blank identifier in Go. It can be used in place of
// any value of any type. Let's see what's the use of this blank identifier.

func main() {
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	// Multiple return values
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)

	// Blank Identifier
	area2, _ := rectProps2(10.8, 5.6) // perimeter is discarded
	fmt.Printf("Area2 %f", area2)
}
