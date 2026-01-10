package main

import "fmt"

// ════════════════════════════════════════════════
// IF-ELSE STATEMENT - Two-Way Conditional Logic
// ════════════════════════════════════════════════
// Executes one block if condition is true, another if false
// Syntax: 
//   if condition { 
//     code when true 
//   } else { 
//     code when false 
//   }

func main() {
	// ════════════════════════════════════════════════
	// CUSTOMER ORDER DATA
	// ════════════════════════════════════════════════
	customerName := "Bogdan"
	orderedSize := "Large"

	// ════════════════════════════════════════════════
	// IF-ELSE LOGIC - Check order size
	// ════════════════════════════════════════════════
	// Compare orderedSize with "Small"
	// == checks for exact equality (case-sensitive)
	
	if orderedSize == "Small" {
		// This runs when: orderedSize == "Small" is TRUE
		fmt.Println(customerName, "ordered a Small coffee. It will be ready in 2 min")
	} else {
		// This runs when: orderedSize == "Small" is FALSE
		// (any other size: Medium, Large, etc.)
		fmt.Println(customerName, "ordered something bigger. It might take a bit more time")
	}
	// Output: "Bogdan ordered something bigger. It might take a bit more time"
	
	// ════════════════════════════════════════════════
	// KEY POINTS
	// ════════════════════════════════════════════════
	// - Only ONE block executes (either if or else)
	// - else is optional, but provides fallback logic
	// - String comparison is case-sensitive: "Small" ≠ "small"
}
