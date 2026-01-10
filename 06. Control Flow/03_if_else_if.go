package main

import "fmt"

// ════════════════════════════════════════════════
// IF-ELSE IF-ELSE STATEMENT - Multi-Way Conditional Logic
// ════════════════════════════════════════════════
// Checks multiple conditions in order, executes first match
// Syntax:
//   if condition1 { code1 }
//   else if condition2 { code2 }
//   else if condition3 { code3 }
//   else { default code }

func main() {
	// ════════════════════════════════════════════════
	// LOYALTY POINTS SYSTEM
	// ════════════════════════════════════════════════
	points := 75

	// ════════════════════════════════════════════════
	// CHAIN OF CONDITIONS - Checked from top to bottom
	// ════════════════════════════════════════════════
	// Stops at FIRST true condition
	
	if points >= 100 {
		// Range: [100, ∞)
		// 75 >= 100 → FALSE, skip to next
		fmt.Println("Platinum member: Free coffee every day!")
		
	} else if points >= 50 {
		// Range: [50, 100)
		// 75 >= 50 → TRUE, executes this block
		fmt.Println("Gold member: 20% discount on latte")
		// ⚠️ Stops here! Remaining conditions are NOT checked
		
	} else if points >= 20 {
		// Range: [20, 50)
		// This is NEVER checked because previous condition matched
		fmt.Println("Silver member: Free cookie on Monday")
		
	} else {
		// Range: [0, 20)
		// Default fallback if no conditions match
		fmt.Println("Bronze member: Keep sipping to earn rewards")
	}
	// Output: "Gold member: 20% discount on latte"

	// ════════════════════════════════════════════════
	// KEY POINTS
	// ════════════════════════════════════════════════
	// - Conditions evaluated TOP to BOTTOM
	// - Only FIRST true condition executes
	// - Order matters! Put specific conditions first
	// - else block is optional (catches all remaining cases)
	// - >= means "greater than or equal to"
}
