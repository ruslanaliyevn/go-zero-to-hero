package main 

import "fmt"

// ============================================
// FUNCTION WITH RETURN VALUE
// ============================================
// calculateLoyaltyPoints calculates loyalty points based on amount spent
// Input: amountSpent (float64) - money customer spent
// Output: int - loyalty points earned (must be whole number)
//              └─── Return type declared here, so return is REQUIRED
func calculateLoyaltyPoints(amountSpent float64) int {
	
	// ============================================
	// CALCULATION + TYPE CONVERSION
	// ============================================
	// Formula: spending × 2 = points
	// amountSpent * 2 gives float64 (e.g., 9.50 * 2 = 19.0)
	// int() converts float64 → int (truncates decimal: 19.0 → 19)
	loyaltyPoints := int(amountSpent * 2)
	
	// Examples:
	// $9.50 → 9.50 * 2 = 19.0 → int(19.0) = 19 points
	// $10.75 → 10.75 * 2 = 21.5 → int(21.5) = 21 points
	
	// ============================================
	// RETURN STATEMENT
	// ============================================
	// Return sends the calculated value back to the caller
	// REQUIRED because function declares return type (int)
	return loyaltyPoints
}

// ============================================
// MAIN FUNCTION
// ============================================
func main() {
	// Call function and capture returned value
	// Function returns 19, which is stored in newlyEarnedPoints
	var newlyEarnedPoints int = calculateLoyaltyPoints(9.50)
	
	fmt.Println("Earned points today:", newlyEarnedPoints)
	// Output: Earned points today: 19
}
