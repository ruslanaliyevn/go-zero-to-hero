package main

import "fmt"

// ════════════════════════════════════════════════
// WITHOUT POINTER - Returns new value
// ════════════════════════════════════════════════
// Receives COPY of price, returns new value
func calculatePriceAfterDiscount(price float64, discountRate float64) float64 {
	// Works with copy - original unchanged
	newPrice := price - (price * discountRate)
	return newPrice  // Must return new value
}

func main() {
	var coffeePrice float64 = 5.00
	var discount float64 = 0.10
	fmt.Printf("Basic coffee price: $%.2f\n", coffeePrice)

	// Must capture returned value
	coffeePrice = calculatePriceAfterDiscount(coffeePrice, discount)
	//            └──────────────┬──────────────┘
	//                    Return value assigned back
	
	fmt.Printf("Price with discount: $%.2f\n", coffeePrice)
	// Output: Price with discount: $4.50
}
