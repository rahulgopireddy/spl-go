package main

import (
	"errors"
	"fmt"
	"time"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second) // Simulate some time-consuming task
	}
}

func main() {
	fmt.Println("Starting the program")

	// Example of Exception Handling
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error in division:", err)
	} else {
		fmt.Println("Result of division:", result)
	}

	// Run printNumbers concurrently using a goroutine
	go func() {
		fmt.Println("Printing numbers concurrently:")
		printNumbers()
	}()

	// Wait for a few seconds to allow goroutine to finish
	time.Sleep(3 * time.Second)

	fmt.Println("End of the program")
}
