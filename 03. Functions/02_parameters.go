package main

import "fmt"

// ============================================
// FUNCTION WITH PARAMETER
// ============================================
// greet accepts one parameter: coffeeShopName (string type)
// Parameters act as input variables for the function
func greet(coffeeShopName string) {
	// ============================================
	// LOCAL VARIABLE SCOPE
	// ============================================
	// greetMessage is a LOCAL variable
	// Created: when function is called
	// Destroyed: when function execution ends
	// Accessible: ONLY inside this function
	var greetMessage string = "Welcome to the Coffee Shop"
	
	// Both greetMessage and coffeeShopName are accessible here
	fmt.Println(greetMessage, coffeeShopName)
	
	// After this closing brace, both variables cease to exist
}

// ============================================
// MAIN FUNCTION - Entry Point
// ============================================
func main() {
	// ============================================
	// FUNCTION CALLS WITH ARGUMENTS
	// ============================================
	// Each call passes a different ARGUMENT to the coffeeShopName PARAMETER
	greet("Brew & Beans")   // "Brew & Beans" is copied to coffeeShopName
	greet("Coffee & Milk")  // "Coffee & Milk" is copied to coffeeShopName
	
	// ============================================
	// SCOPE ERRORS - Variables Don't Exist Here
	// ============================================
	
	// ❌ ERROR: coffeeShopName is undefined here
	// fmt.Println(coffeeShopName)
	// Reason: coffeeShopName is a PARAMETER of greet() function
	//         Parameters only exist inside their function
	//         main() cannot see greet()'s parameters
	
	// ❌ ERROR: greetMessage is undefined here
	// fmt.Println(greetMessage)
	// Reason: greetMessage is a LOCAL variable inside greet() function
	//         Local variables only exist inside their function
	//         main() cannot see greet()'s local variables
	
	// KEY CONCEPT: Variables have LIMITED SCOPE
	// They only exist within the function where they're declared
}
