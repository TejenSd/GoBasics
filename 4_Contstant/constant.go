package main

import "fmt"

func main() {
	// Constant declaration
	const pi = 3.14
	const greeting = "Hello, World!"

	// Printing constants
	fmt.Println("Value of pi:", pi)
	fmt.Println(greeting)

	// Using constants in calculations
	radius := 5.0
	area := pi * radius * radius
	fmt.Println("Area of circle with radius", radius, "is:", area)

	const (
		maxUsers = 100
		minUsers = 1
	)
	fmt.Println("Max users allowed:", maxUsers)
	fmt.Println("Min users allowed:", minUsers)
}
