package main

import "fmt"

func main() {
	// ════════════════════════════════════════════════
	// ARRAY DECLARATION
	// ════════════════════════════════════════════════
	// Declare array with size 3 - all elements default to ""
	var coffeeSizes [3]string
	fmt.Println(coffeeSizes)
	// Output: [  ] (three empty strings)

	// ════════════════════════════════════════════════
	// ACCESSING ARRAY ELEMENTS
	// ════════════════════════════════════════════════
	// Arrays use zero-based indexing: [0], [1], [2]
	coffeeSizes[0] = "Small"
	fmt.Println(coffeeSizes)
	// Output: [Small  ]

	coffeeSizes[1] = "Medium"
	coffeeSizes[2] = "Large"
	fmt.Println(coffeeSizes)
	// Output: [Small Medium Large]

	// ════════════════════════════════════════════════
	// MODIFYING ARRAY ELEMENTS
	// ════════════════════════════════════════════════
	// Can update existing elements
	coffeeSizes[2] = "Extra Large"
	fmt.Println(coffeeSizes)
	// Output: [Small Medium Extra Large]

	// ════════════════════════════════════════════════
	// ARRAY BOUNDS
	// ════════════════════════════════════════════════
	// ❌ This would cause ERROR - index out of bounds
	// coffeeSizes[4] = "Super Extra Large"
	// Error: invalid argument: index 4 out of bounds [0:3]
	// Array size is 3, so valid indices are: 0, 1, 2

	// ════════════════════════════════════════════════
	// READING ARRAY ELEMENTS
	// ════════════════════════════════════════════════
	fmt.Println("First element is:", coffeeSizes[0])
	// Output: First element is: Small
}
