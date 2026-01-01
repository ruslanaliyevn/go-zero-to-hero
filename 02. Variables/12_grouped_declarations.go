package main

import "fmt"

func main() {
	// ============================================
	// INDIVIDUAL VARIABLE DECLARATIONS
	// ============================================
	// Order details declared individually
	var coffeeType string = "Latte"
	var quantity int = 3
	var unitPrice float64 = 4.25

	// Print order information
	// %d = integer, %s = string, %.2f = float with 2 decimal places
	fmt.Printf("Ordered %d %s priced at $%.2f each\n", quantity, coffeeType, unitPrice)

	// ============================================
	// GROUPED VARIABLE DECLARATIONS
	// ============================================
	// Customer information grouped together for better organization
	var (
		customerName string = "Bogdan"
		tableNum     int    = 8
		isReadyToPay bool   = false
	)

	// Print customer information
	// %t = boolean (true/false)
	fmt.Printf("Customer %s at table %d is ready to pay: %t\n", customerName, tableNum, isReadyToPay)

	// ============================================
	// GROUPED CONSTANT DECLARATIONS
	// ============================================
	// Coffee size constants - these values never change
	const (
		sizeSmall  = "S"
		sizeMedium = "M"
		sizeLarge  = "L"
	)

	// ============================================
	// KEY DIFFERENCE: var vs const
	// ============================================
	// ⚠️ UNUSED VARIABLES → ❌ Compilation error in Go
	// Example: var unused = "test" → ERROR: unused declared and not used
	
	// ✅ UNUSED CONSTANTS → ✅ No error
	// The size constants above are unused, but no compilation error occurs
	// This is a key difference between var and const in Go
}
