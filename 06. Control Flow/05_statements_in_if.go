package main

import "fmt"

// ════════════════════════════════════════════════
// HELPER FUNCTION - Calculate order with tax
// ════════════════════════════════════════════════
func getOrderWithTax(amount float64, tax float64) float64 {
	return amount + amount*tax
}

// ════════════════════════════════════════════════
// IF STATEMENT WITH INITIALIZATION
// ════════════════════════════════════════════════
// Go allows variable declaration/initialization in if statement
// Syntax: if statement; condition { code }
// Scope: Variable only exists within if-else block

func main() {
	// ════════════════════════════════════════════════
	// EXAMPLE 1: Simple initialization
	// ════════════════════════════════════════════════
	// Declare 'points' and immediately use it in condition
	// points := 15 → assignment
	// points > 10  → condition check
	if points := 15; points > 10 {
		// 15 > 10 → TRUE
		fmt.Println("You are eligible for coffee discount")
	}
	// ⚠️ 'points' variable is NOT accessible here (out of scope)

	// ════════════════════════════════════════════════
	// EXAMPLE 2: Initialization with function call
	// ════════════════════════════════════════════════
	// Calculate total: 14.50 + 14.50*0.1 = 14.50 + 1.45 = 15.95
	// fullAmount := 15.95
	// fullAmount > 15 → TRUE
	if fullAmount := getOrderWithTax(14.50, 0.1); fullAmount > 15 {
		fmt.Println("You can join coffee club")
	}
	// ⚠️ 'fullAmount' is NOT accessible here (out of scope)

	// ════════════════════════════════════════════════
	// EXAMPLE 3: Increment operator (++)
	// ════════════════════════════════════════════════
	totalLoyaltyPoints := 150
	
	// totalLoyaltyPoints++ → increment first (150 → 151)
	// totalLoyaltyPoints > 120 → 151 > 120 → TRUE
	if totalLoyaltyPoints++; totalLoyaltyPoints > 120 {
		// totalLoyaltyPoints is now 151
		fmt.Println("Total loyalty points:", totalLoyaltyPoints) // 151
		fmt.Println("You can become Gold member")
	}

	// ════════════════════════════════════════════════
	// EXAMPLE 4: Compound assignment (+=)
	// ════════════════════════════════════════════════
	// totalLoyaltyPoints += 10 → 151 + 10 = 161
	// totalLoyaltyPoints > 120 → 161 > 120 → TRUE
	if totalLoyaltyPoints += 10; totalLoyaltyPoints > 120 {
		// totalLoyaltyPoints is now 161
		fmt.Println("Total loyalty points:", totalLoyaltyPoints) // 161
		fmt.Println("You can become Gold member")
	}

	// ════════════════════════════════════════════════
	// KEY BENEFITS
	// ════════════════════════════════════════════════
	// ✓ Limits variable scope (prevents pollution)
	// ✓ Cleaner code (declaration + condition in one line)
	// ✓ Variables are temporary and don't leak outside
	
	// ════════════════════════════════════════════════
	// STATEMENT TYPES ALLOWED
	// ════════════════════════════════════════════════
	// - Variable declaration: x := 10
	// - Function call: result := func()
	// - Increment/Decrement: x++, x--
	// - Assignments: x += 5, x = y * 2
}
