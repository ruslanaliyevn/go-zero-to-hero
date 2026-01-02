package main

import "fmt"

func main() {
	// ============================================
	// OUTER VARIABLE (for closure)
	// ============================================
	// Tax rate used in the anonymous function
	taxRate := 0.10  // 10% tax
	
	// ============================================
	// FUNCTION LITERAL (Anonymous Function)
	// ============================================
	// Creating a function without a name and assigning it to a variable
	// This is called an "anonymous function" or "function literal"
	calculateTax := func(amount float64) float64 {
		//             └──────────┬──────────┘
		//                 No function name here
		//                 Assigned directly to calculateTax variable
		
		// ============================================
		// CLOSURE: Access to outer variable
		// ============================================
		// The function can "see" and use taxRate from outer scope
		// This is called a "closure"
		return amount * taxRate
		//              └──┬──┘
		//          Uses outer variable (closure feature)
	}
	
	// ============================================
	// USING THE FUNCTION LITERAL
	// ============================================
	// Call it like a regular function using the variable name
	subtotal := 25.00
	tax := calculateTax(subtotal)
	//     └──────┬──────┘
	//     Called like a normal function
	//     25.00 × 0.10 = 2.50
	
	total := subtotal + tax  // 25.00 + 2.50 = 27.50
	
	fmt.Printf("Total amount to pay: $%.2f\n", total)
	// Output: Total amount to pay: $27.50
}
