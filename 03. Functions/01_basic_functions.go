package main

import "fmt"

// ============================================
// FUNCTION DEFINITION
// ============================================
// greet is a simple function with no parameters and no return value
// func = keyword to define a function
// greet = function name (should be descriptive)
// () = empty parentheses means no parameters
// {} = function body contains the code to execute
func greet() {
	fmt.Println("Welcome to the Coffee Shop", "Brew & Beans")
}

// ============================================
// MAIN FUNCTION - Entry Point
// ============================================
// main() is the entry point of every Go program
// Execution starts here
func main() {
	// Call the greet function - executes the code inside greet()
	greet()
	
	// Function can be called multiple times
	greet()
	
	// Output:
	// Welcome to the Coffee Shop Brew & Beans
	// Welcome to the Coffee Shop Brew & Beans
}
