package main

import "fmt"

func main() {
	// Array (Slice in Go)
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range numbers {
		sum += num // Calculating sum using a for loop
	}
	fmt.Println("Sum of numbers in the array:", sum)

	// Struct (Object in Go)
	type Person struct {
		Name      string
		Age       int
		IsEmployed bool
	}

	person := Person{
		Name:      "Rahul",
		Age:       25,
		IsEmployed: true,
	}

	if person.IsEmployed {
		fmt.Println(person.Name, "is a Student.")
	} else {
		fmt.Println(person.Name, "is not a Student.")
	}

	// Control Structures
	num := 10
	if num > 0 {
		fmt.Println(num, "is a positive number.")
	} else if num < 0 {
		fmt.Println(num, "is a negative number.")
	} else {
		fmt.Println(num, "is zero.")
	}

	// Lambda function (Anonymous function in Go)
	square := func(x int) int {
		return x * x
	}
	fmt.Println("Square of 5:", square(5))
}
