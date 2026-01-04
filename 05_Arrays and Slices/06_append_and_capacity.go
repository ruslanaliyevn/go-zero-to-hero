package main

import "fmt"

func main() {
	// ════════════════════════════════════════════════
	// INITIAL SLICE
	// ════════════════════════════════════════════════
	// Create slice with 2 elements
	menu := []string{"Cake", "Pie"}

	fmt.Println("Initial menu:", menu)
	fmt.Println("Length:", len(menu), "Capacity:", cap(menu))
	// Length: 2 (number of elements)
	// Capacity: 2 (underlying array size)
	fmt.Printf("Memory location: %p\n", &menu)
	fmt.Printf("Memory location of \"Cake\": %p\n", &menu[0])
	fmt.Println("-------------------")

	// ════════════════════════════════════════════════
	// FIRST APPEND - Capacity Increases
	// ════════════════════════════════════════════════
	// Adding element when capacity is full
	// Go creates NEW underlying array with DOUBLE capacity
	menu = append(menu, "Donut")
	fmt.Println("Menu after adding donut:", menu)
	fmt.Println("Length:", len(menu), "Capacity:", cap(menu))
	// Length: 3, Capacity: 4 (doubled from 2)
	// Notice: Memory location of "Cake" CHANGED - new array created!
	fmt.Printf("Memory location: %p\n", &menu)
	fmt.Printf("Memory location of \"Cake\": %p\n", &menu[0])
	fmt.Println("-------------------")

	// ════════════════════════════════════════════════
	// SECOND APPEND - Within Capacity
	// ════════════════════════════════════════════════
	// Capacity is 4, length is 3 - still has room
	menu = append(menu, "Ice Cream")
	fmt.Println("Menu after adding ice cream:", menu)
	fmt.Println("Length:", len(menu), "Capacity:", cap(menu))
	// Length: 4, Capacity: 4 (no change)
	// Memory location stays SAME - no new array needed
	fmt.Printf("Memory location: %p\n", &menu)
	fmt.Printf("Memory location of \"Cake\": %p\n", &menu[0])
	fmt.Println("-------------------")

	// ════════════════════════════════════════════════
	// THIRD APPEND - Capacity Full Again
	// ════════════════════════════════════════════════
	// Capacity is 4, length is 4 - no room left
	// Go creates NEW array with capacity 8 (doubled)
	menu = append(menu, "Cream")
	fmt.Println("Menu after adding cream", menu)
	fmt.Println("Length:", len(menu), "Capacity:", cap(menu))
	// Length: 5, Capacity: 8 (doubled from 4)
	// Memory location CHANGED again - new array created!
	fmt.Printf("Memory location: %p\n", &menu)
	fmt.Printf("Memory location of \"Cake\": %p\n", &menu[0])
	fmt.Println("-------------------")

	// ════════════════════════════════════════════════
	// MAKE WITH LENGTH AND CAPACITY
	// ════════════════════════════════════════════════
	fmt.Println("-------------------")
	// make([]Type, length, capacity)
	// Creates slice with length 0 but capacity 5
	cupSizes := make([]string, 0, 5)
	//                         └┬┘ └┬┘
	//                       len  cap
	
	fmt.Println("Len of cupSizes:", len(cupSizes), "Capacity of cupSizes:", cap(cupSizes))
	// Len: 0, Capacity: 5

	// ════════════════════════════════════════════════
	// LENGTH vs CAPACITY
	// ════════════════════════════════════════════════
	// ❌ Can't assign to index 0 - length is 0!
	// cupSizes[0] = "Small"
	// Error: panic: runtime error: index out of range [0] with length 0
	
	// ✅ Must use append() when length is 0
	cupSizes = append(cupSizes, "Small", "Medium")
	// Now length becomes 2
	fmt.Println("Len of cupSizes:", len(cupSizes), "Capacity of cupSizes:", cap(cupSizes))
	// Len: 2, Capacity: 5

	// ✅ NOW we can assign - length is 2, so index 0 and 1 are valid
	cupSizes[0] = "Extra small"
	fmt.Println(cupSizes)
	// Output: [Extra small Medium]
	fmt.Println("Len of cupSizes:", len(cupSizes), "Capacity of cupSizes:", cap(cupSizes))
	// Len: 2, Capacity: 5
}
