package main

import "fmt"

func main() {
	// ════════════════════════════════════════════════
	// STEP 1: Create original variable
	// ════════════════════════════════════════════════
	var coffeePrice = 4.50
	fmt.Println("Coffee price", coffeePrice)
	// STEP 1
	// Compile Time (code you write): var coffeePrice = 4.50
	// Runtime (what machine sees):  [some memory address] = 4.50

	// ════════════════════════════════════════════════
	// STEP 2: Display memory address
	// ════════════════════════════════════════════════
	// STEP 2
	// Compile Time (code you write):	fmt.Println("Coffee price", coffeePrice)
	// Runtime (what machine sees):  fmt.Println([some mem address], [memory address (same as in step 1)] )
	fmt.Println("Memory address of price 5.00", &coffeePrice)
	// Output: Memory address of price 5.00 0xc00000a088
	// & = address-of operator

	// ════════════════════════════════════════════════
	// STEP 3: Update value directly
	// ════════════════════════════════════════════════
	coffeePrice = 5.00
	fmt.Println("Updated coffee of price 5.00", &coffeePrice)
	// Output: Updated coffee of price 5.00 0xc00000a088
	// IMPORTANT: Same memory address, just value changed

	// ════════════════════════════════════════════════
	// STEP 4: Create pointer variable
	// ════════════════════════════════════════════════
	// pointerToCoffeePrice := &coffeePrice // same as next line
	var pointerToCoffeePrice *float64 = &coffeePrice
	//                       └───┬───┘   └─────┬─────┘
	//                      Type: pointer    Get address
	//                      to float64

	// ════════════════════════════════════════════════
	// STEP 5: Change value THROUGH pointer
	// ════════════════════════════════════════════════
	/* got to the memory location where pointerToCoffeePrice points to
	and change value in this memory location */
	*pointerToCoffeePrice = 7.50
	// └──────┬──────┘
	//    * = dereference operator
	//    Go to address stored in pointerToCoffeePrice
	//    and change the value there to 7.50

	// ════════════════════════════════════════════════
	// STEP 6: Verify - original variable changed!
	// ════════════════════════════════════════════════
	fmt.Println("Updated coffeePrice value in memory ", coffeePrice)
	// Output: Updated coffeePrice value in memory 7.5
	// 
	// KEY: coffeePrice is now 7.5 because we changed it via pointer
	// Both point to SAME memory location (0xc00000a088)
}

