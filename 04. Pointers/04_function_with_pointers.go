package main

import "fmt"

// ════════════════════════════════════════════════
// WITH POINTER - Modifies original directly
// ════════════════════════════════════════════════
// Receives ADDRESS of price, no return needed
func applyDiscount(price *float64, discountRate float64) {
	//                └────┬────┘
	//              Pointer to float64
	
	// Dereference (*) to modify value at address
	*price = *price - (*price * discountRate)
	// └┬┘   └──┬──┘   └───┬───┘
	//  │      │           └─ Read value at address
	//  │      └─ Read value at address
	//  └─ Write new value to address
	
	// No return - original is modified directly
}

func main() {
	var coffeePrice float64 = 5.00
	var discount float64 = 0.10
	fmt.Printf("Basic coffee price: $%.2f\n", coffeePrice)

	// Pass address with &
	applyDiscount(&coffeePrice, discount)
	//            └─────┬─────┘
	//            Send address, not copy
	
	// coffeePrice automatically changed!
	fmt.Printf("Price with discount: $%.2f\n", coffeePrice)
	// Output: Price with discount: $4.50
}
