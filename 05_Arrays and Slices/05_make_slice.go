package main

import "fmt"

func main() {
	// ════════════════════════════════════════════════
	// SLICE LITERAL - Direct initialization
	// ════════════════════════════════════════════════
	// Create slice with initial values
	// Note: No size specified in [] - this makes it a SLICE, not array
	ratings := []int{5, 4, 5, 5, 3}
	//         └┬┘
	//   Empty brackets = slice (dynamic size)
	//   vs [5]int = array (fixed size)
	
	fmt.Println("Original ratings", ratings)
	// Output: Original ratings [5 4 5 5 3]

	// ════════════════════════════════════════════════
	// MODIFY SLICE ELEMENT
	// ════════════════════════════════════════════════
	// Change element at index 2
	ratings[2] = 3
	fmt.Println("Changed element with index 2:", ratings)
	// Output: Changed element with index 2: [5 4 3 5 3]

	// ════════════════════════════════════════════════
	// SLICE LENGTH
	// ════════════════════════════════════════════════
	fmt.Println("Length of the slice is:", len(ratings))
	// Output: Length of the slice is: 5

	// ════════════════════════════════════════════════
	// CREATE SLICE WITH make()
	// ════════════════════════════════════════════════
	// make([]Type, length) creates slice with specified length
	// All elements initialized to zero values ("" for strings)
	coffeeTypes := make([]string, 3)
	//             └─────┬─────┘  └┬┘
	//              make function   │
	//                          length = 3
	
	fmt.Println("Coffee types after make:", coffeeTypes)
	// Output: Coffee types after make: [  ]
	//         (three empty strings)

	// ════════════════════════════════════════════════
	// ASSIGN VALUES TO SLICE
	// ════════════════════════════════════════════════
	// Fill the slice with values
	coffeeTypes[0] = "Cappuccino"
	coffeeTypes[1] = "Latte"
	coffeeTypes[2] = "Espresso"

	fmt.Println("Coffee types after reassignment:", coffeeTypes)
	// Output: Coffee types after reassignment: [Cappuccino Latte Espresso]
}
