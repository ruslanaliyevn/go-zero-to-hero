
package main

import "fmt"

func main() {
	// ════════════════════════════════════════════════
	// ARRAY DECLARATION
	// ════════════════════════════════════════════════
	// Fixed-size array with 4 elements
	dessertMenu := [4]string{"Muffin", "Brownie", "Croissant", "Cookie"}
	//                       index: 0       1         2          3
	fmt.Println("Dessert Menu:", dessertMenu)
	// Output: Dessert Menu: [Muffin Brownie Croissant Cookie]

	// ════════════════════════════════════════════════
	// SLICE - Subset of Array
	// ════════════════════════════════════════════════
	// Syntax: array[start:end]
	// - Includes start index
	// - Excludes end index
	// - Creates a "view" into the original array

	// ════════════════════════════════════════════════
	// RANGE SLICE [1:3]
	// ════════════════════════════════════════════════
	// Start at index 1, stop before index 3
	// Gets elements at indices: 1, 2
	slice := dessertMenu[1:3]
	fmt.Println("Slice of the Dessert Menu [1:3]", slice)
	// Output: [Brownie Croissant]

	// ════════════════════════════════════════════════
	// ALL ELEMENTS [:]
	// ════════════════════════════════════════════════
	// Empty start and end = get all elements
	slice = dessertMenu[:]
	fmt.Println("Slice of the Dessert Menu [:]", slice)
	// Output: [Muffin Brownie Croissant Cookie]

	// ════════════════════════════════════════════════
	// FROM INDEX TO END [2:]
	// ════════════════════════════════════════════════
	// Start at index 2, go to end
	// Gets elements at indices: 2, 3
	slice = dessertMenu[2:]
	fmt.Println("Slice of the Dessert Menu [2:]", slice)
	// Output: [Croissant Cookie]

	// ════════════════════════════════════════════════
	// FROM START TO INDEX [:3]
	// ════════════════════════════════════════════════
	// Start from beginning, stop before index 3
	// Gets elements at indices: 0, 1, 2
	slice = dessertMenu[:3]
	fmt.Println("Slice of the Dessert Menu [:3]", slice)
	// Output: [Muffin Brownie Croissant]
}
