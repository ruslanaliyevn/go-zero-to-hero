package main

import "fmt"

// ════════════════════════════════════════════════
// MULTIPLE CONDITIONS - Logical Operators
// ════════════════════════════════════════════════
// Combine multiple conditions using logical operators:
//   && (AND) - ALL conditions must be true
//   || (OR)  - AT LEAST ONE condition must be true
//   !  (NOT) - Inverts boolean value

func main() {
	// ════════════════════════════════════════════════
	// HAPPY HOUR PROMOTION RULES
	// ════════════════════════════════════════════════
	// Requirements (ALL must be met):
	//   1. Time: 15:00 - 17:00 (3 PM - 5 PM)
	//   2. Customer must be a member
	//   3. Order amount must exceed $10
	
	hour := 16
	isMember := true
	orderAmount := 13.50

	// ════════════════════════════════════════════════
	// AND (&&) OPERATOR - ALL conditions must be TRUE
	// ════════════════════════════════════════════════
	// Evaluation:
	//   hour >= 15        → 16 >= 15  → TRUE  ✓
	//   hour <= 17        → 16 <= 17  → TRUE  ✓
	//   isMember          → true      → TRUE  ✓
	//   orderAmount > 10  → 13.50 > 10 → TRUE  ✓
	//   
	// Result: TRUE && TRUE && TRUE && TRUE = TRUE
	
	if hour >= 15 && hour <= 17 && isMember && orderAmount > 10 {
		// All conditions met - customer qualifies
		fmt.Println("You get 30% off!")
	} else {
		// At least one condition failed
		fmt.Println("No happy hour deals available")
	}
	// Output: "You get 30% off!"

	// ════════════════════════════════════════════════
	// LOGICAL OPERATORS OVERVIEW
	// ════════════════════════════════════════════════
	// AND (&&): true && true = true
	//           true && false = false
	//           false && false = false
	//
	// OR (||):  true || false = true
	//           false || false = false
	//
	// NOT (!):  !true = false
	//           !false = true
	
	// ════════════════════════════════════════════════
	// SHORT-CIRCUIT EVALUATION
	// ════════════════════════════════════════════════
	// && stops checking if first condition is false
	// || stops checking if first condition is true
}
