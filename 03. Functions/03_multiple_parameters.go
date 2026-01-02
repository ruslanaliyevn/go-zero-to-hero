package main

import "fmt"

// ============================================
// FUNCTION WITH MULTIPLE PARAMETERS
// ============================================
// Three parameters: string, string, float64
// Order of arguments must match parameter order
func getDrinkInfo(CustomerName string, drink string, price float64) {
	
	// ============================================
	// TWO WAYS TO USE Printf
	// ============================================
	
	// METHOD 1: Direct format string (inline)
	// fmt.Printf("%s's favorite drink is %s and it's price is $%.2f\n", CustomerName, drink, price)
	
	// METHOD 2: Stored format string (reusable, cleaner)
	info := "%s's favorite drink is %s and it's price is $%.2f\n"
	fmt.Printf(info, CustomerName, drink, price)
	
	// Stored format is better for:
	// ✅ Reusability - use same format multiple times
	// ✅ Readability - separates format from data
	// ✅ Maintainability - change format in one place
}

// ============================================
// MAIN FUNCTION
// ============================================
func main() {
	// Call function with different arguments
	// Arguments must match: (string, string, float64)
	getDrinkInfo("Bogdan", "Capuccino", 4.50)
	getDrinkInfo("Alice", "Espresso", 2.50)
}
