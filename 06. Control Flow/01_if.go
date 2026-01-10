package main

import "fmt"

// ════════════════════════════════════════════════
// IF STATEMENT - Conditional Logic
// ════════════════════════════════════════════════
// Executes code block only when condition is true
// Syntax: if condition { code }

func main() {
	// ════════════════════════════════════════════════
	// SCENARIO 1: Condition is TRUE
	// ════════════════════════════════════════════════
	OrderTotal := 15.0
	
	// Check if order qualifies for free cookie
	// 15.0 > 10 → TRUE
	if OrderTotal > 10 {
		fmt.Println("You get a free cookie with your order!")
	}
	// Output: "You get a free cookie with your order!"

	// ════════════════════════════════════════════════
	// SCENARIO 2: Condition is FALSE
	// ════════════════════════════════════════════════
	OrderTotal = 7.50
	
	// Check again with new value
	// 7.50 > 10 → FALSE
	if OrderTotal > 10 {
		fmt.Println("You get a free cookie with your order!")
	}
	// Output: (nothing - condition not met)
	
	// ════════════════════════════════════════════════
	// COMPARISON OPERATORS
	// ════════════════════════════════════════════════
	// >  greater than
	// <  less than
	// >= greater than or equal
	// <= less than or equal
	// == equal to
	// != not equal to
}
