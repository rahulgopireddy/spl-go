package main

import (
	"fmt"
	"math"
)

func main() {
	// Integer
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	average := float64(num1+num2) / 2 // Average
	fmt.Println("Average:", average)

	// Float
	float1 := 3.14
	float2 := 2.71
	product := float1 * float2
	fmt.Println("Product:", product)

	roundedValue := math.Round(product) // Rounding to the nearest integer
	fmt.Println("Rounded Value:", roundedValue)

	// String
	str1 := "Hello, "
	str2 := "World!"
	combinedString := str1 + str2 // Concatenation
	fmt.Println("Combined String:", combinedString)

	replacedString := replaceString(combinedString, "World", "Golang") // String replacement
	fmt.Println("Replaced String:", replacedString)

	// Boolean
	isTrue := true
	isFalse := false
	isAnyTrue := isTrue || isFalse // Logical OR
	fmt.Println("Is any true?", isAnyTrue)

	isBothTrue := isTrue && isFalse // Logical AND
	fmt.Println("Are both true?", isBothTrue)
}

func replaceString(input, old, new string) string {
	return input[:len(input)-len(old)] + new
}
