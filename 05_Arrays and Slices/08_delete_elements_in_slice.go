package main

import "fmt"

// ════════════════════════════════════════════════
// DELETE BY INDEX - Reusable Function
// ════════════════════════════════════════════════
// Removes element at specified index from slice
// Parameters:
//   - index: position to remove
//   - slice: original slice
// Returns: new slice without the element
func deleteByIndex(index int, slice []string) []string {
	// Combine elements before and after the index
	// slice[:index] - everything before index
	// slice[index+1:] - everything after index
	// ... unpacks the second slice into individual elements
	return append(slice[:index], slice[index+1:]...)
	//            └─────┬─────┘  └──────┬──────┘
	//          Before index    After index
}

func main() {
	// ════════════════════════════════════════════════
	// INITIAL SLICE
	// ════════════════════════════════════════════════
	coffees := []string{"Espresso", "Latte", "Mocha", "Cappuccino"}
	//        indices:    0          1        2        3

	fmt.Println("Original menu:", coffees)
	fmt.Println("Length is:", len(coffees), "Capacity is:", cap(coffees))
	// Length: 4, Capacity: 4

	// ════════════════════════════════════════════════
	// DELETE INLINE - Remove "Latte" at index 1
	// ════════════════════════════════════════════════
	indexToRemove := 1
	
	// How it works:
	// coffees[:1] = [Espresso]
	// coffees[2:] = [Mocha, Cappuccino]
	// append([Espresso], [Mocha, Cappuccino]...)
	// Result: [Espresso, Mocha, Cappuccino]
	coffees = append(coffees[:indexToRemove], coffees[indexToRemove+1:]...)
	
	fmt.Println("Updated menu without Latte:", coffees)
	fmt.Println("Length is:", len(coffees), "Capacity is:", cap(coffees))
	// Length: 3, Capacity: 4 (capacity unchanged)

	// ════════════════════════════════════════════════
	// DELETE WITH FUNCTION - Remove "Espresso" at index 0
	// ════════════════════════════════════════════════
	indexToRemove = 0
	
	// Using the reusable function
	coffees = deleteByIndex(0, coffees)
	
	fmt.Println("Updated menu without Espresso:", coffees)
	fmt.Println("Length is:", len(coffees), "Capacity is:", cap(coffees))
	// Length: 2, Capacity: 4 (capacity still unchanged)
	// Result: [Mocha, Cappuccino]
}
