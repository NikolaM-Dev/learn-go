package main

import "fmt"

type employee struct {
	salary  int
	country string
}

func main() {
	employeeSalary := make(map[string]int)
	employeeSalary["steve"] = 12000
	employeeSalary["jaime"] = 15000
	employeeSalary["mike"] = 9000
	fmt.Println("employeeSalary map contents:", employeeSalary)

	// It is also possible to initialize a map during the declaration itself.
	employeeSalary2 := map[string]int{
		"steve": 12000,
		"jaime": 15000,
	}
	employeeSalary2["mike"] = 9000
	fmt.Println("employeeSalary2 map contents:", employeeSalary2)

	// Zero value of a map
	// The zero value of a map is nil. If you try to add elements to a nil map,
	// a run-time panic will occur. Hence the map has to be initialized before adding
	// elements.
	/*
		var employeeSalary3 map[string]int
		employeeSalary3["steve"] = 12000 // panic: assignment to entry in nil map
	*/

	// Retrieving value for a key from a map
	employee1 := "jaime"
	salary := employeeSalary[employee1]
	fmt.Println("Salary of", employee1, "is", salary)

	// What will happen if an element is not present? The map will return the zero
	// value of the type of that element. In the case of employeeSalary map, if we
	// try to access an element which is not present, the zero value of int which
	// is 0 will be returned.
	fmt.Println("Salary of joe is", employeeSalary["joe"]) // Salary of joe is 0

	// Checking if a key exists
	newEmp := "joe"
	value, ok := employeeSalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
		return
	}
	fmt.Println(newEmp, "not found")

	// Iterate over all elements in a map
	fmt.Println("Contents of the map")
	for key, value := range employeeSalary {
		fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	}

	// Deleting items from a map
	fmt.Println("map before deletion", employeeSalary)
	delete(employeeSalary, "steve")
	fmt.Println("map after deletion", employeeSalary)
	// If we try to delete a key that is not present in the map, there will
	// be no runtime error.

	// Map of structs
	emp1 := employee{
		salary:  12000,
		country: "USA",
	}
	emp2 := employee{
		salary:  14000,
		country: "Canada",
	}
	emp3 := employee{
		salary:  13000,
		country: "India",
	}

	employeeInfo := map[string]employee{
		"Steve": emp1,
		"Jamie": emp2,
		"Mike":  emp3,
	}

	for name, info := range employeeInfo {
		fmt.Printf("Employee: %s Salary: $%d Country: %s\n", name, info.salary, info.country)
	}

	// Length of the map
	fmt.Println("length is", len(employeeSalary))

	// Maps are reference types
	// Similar to slices, maps are reference types
	fmt.Println("Original employee salary", employeeSalary2)
	modified := employeeSalary2
	modified["mike"] = 18000
}
