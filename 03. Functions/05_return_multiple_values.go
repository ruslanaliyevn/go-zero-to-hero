package main 

import "fmt"

// ============================================
// FUNCTION WITH MULTIPLE RETURN VALUES
// ============================================
// processPayment calculates total amount due and change
// Input: orderTotal (float64) - cost of order
//        tip (float64) - tip amount
//        amountPaid (float64) - money customer gave
// Output: (float64, float64) - returns TWO values:
//         1st: total amount due (order + tip)
//         2nd: change to return to customer
func processPayment(orderTotal float64, tip float64, amountPaid float64) (float64, float64) {
	//                                                                      └─────┬─────┘
	//                                              Two return types in parentheses
	
	// ============================================
	// CALCULATE TOTAL AMOUNT DUE
	// ============================================
	totalAmountDue := orderTotal + tip
	// Example: $6.50 + $2.00 = $8.50
	
	// ============================================
	// CALCULATE CHANGE
	// ============================================
	change := amountPaid - totalAmountDue
	// Example: $10.00 - $8.50 = $1.50
	
	// ============================================
	// RETURN MULTIPLE VALUES
	// ============================================
	// Return both values separated by comma
	return totalAmountDue, change
	//     └──────┬──────┘  └──┬─┘
	//            │             └─── 2nd return value
	//            └───────────────── 1st return value
}

// ============================================
// MAIN FUNCTION
// ============================================
func main() {
	// ════════════════════════════════════════════════
	// TRANSACTION 1: Receiving multiple return values
	// ════════════════════════════════════════════════
	// Use multiple variables to capture both returned values
	totalCost, change := processPayment(6.50, 2.00, 10.00)
	//        └─────┬─────┘            └───────┬───────┘
	//              │                          └─── Arguments passed in
	//              └─── Two variables to receive two return values
	
	// totalCost receives 1st return value (8.50)
	// change receives 2nd return value (1.50)
	
	fmt.Printf("Total cost (with tip): $%.2f\n", totalCost)
	fmt.Printf("Change returned to the customer: $%.2f\n", change)
	// Output:
	// Total cost (with tip): $8.50
	// Change returned to the customer: $1.50
	
	fmt.Println("_____________________")
	
	// ════════════════════════════════════════════════
	// TRANSACTION 2: Reusing same variables
	// ════════════════════════════════════════════════
	// Can reuse existing variables with = (not :=)
	totalCost, change = processPayment(28.50, 1.50, 50.00)
	//                ▲ Note: = not := (variables already exist)
	
	fmt.Printf("Total cost (with tip): $%.2f\n", totalCost)
	fmt.Printf("Change returned to the customer: $%.2f\n", change)
	// Output:
	// Total cost (with tip): $30.00
	// Change returned to the customer: $20.00
}
