package main

import "fmt"

// ════════════════════════════════════════════════
// RANGE - Iterating over a slice
// ════════════════════════════════════════════════
// The 'range' keyword iterates over each element in a slice
// Each iteration returns 2 values:
//   1. index - position of the element (starts from 0)
//   2. value - the element itself

func main() {
	// ════════════════════════════════════════════════
	// INITIAL SLICE
	// ════════════════════════════════════════════════
	menu := []string{"Espresso", "Latte", "Cake", "Ice Cream", "Chocolate"}
	//      indices:    0          1        2        3            4

	fmt.Println("Today's menu:")

	// ════════════════════════════════════════════════
	// RANGE LOOP - Iterate through the slice
	// ════════════════════════════════════════════════
	// index: 0, 1, 2, 3, 4
	// menuItem: value of each element in the slice
	for index, menuItem := range menu {
		// index+1 because we want to display starting from 1 for users
		// %d - number (decimal)
		// %s - string
		fmt.Printf("%d. %s\n", index+1, menuItem)
		//          │   │
		//          │   └─> menuItem value
		//          └─────> index + 1
	}
}
