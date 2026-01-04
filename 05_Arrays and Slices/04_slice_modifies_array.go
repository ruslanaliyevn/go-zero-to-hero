package main

import "fmt"

func main() {
	// ════════════════════════════════════════════════
	// ARRAY AND SLICE CREATION
	// ════════════════════════════════════════════════
	// Original array with 3 elements
	menu := [3]string{"Tea", "Coffee", "Juice"}
	//       indices:    0      1        2
	
	// Create slice from array - first 2 elements
	// IMPORTANT: Slice is a VIEW into the array, NOT a copy!
	slice := menu[:2]
	//       └───┬───┘
	//      Gets indices 0, 1

	// ════════════════════════════════════════════════
	// BEFORE MODIFICATION
	// ════════════════════════════════════════════════
	fmt.Println("Before slice change:")
	fmt.Println("menu", menu)
	// Output: menu [Tea Coffee Juice]
	fmt.Println("slice:", slice)
	// Output: slice: [Tea Coffee]

	// ════════════════════════════════════════════════
	// LENGTH COMPARISON
	// ════════════════════════════════════════════════
	fmt.Println("Length of menu array:", len(menu))
	// Output: Length of menu array: 3
	fmt.Println("Length of slice:", len(slice))
	// Output: Length of slice: 2

	// ════════════════════════════════════════════════
	// MODIFY SLICE - THIS AFFECTS THE ORIGINAL ARRAY!
	// ════════════════════════════════════════════════
	// Change first element in the slice
	slice[0] = "Matcha"
	// Since slice references the array, this changes menu[0] too!

	// ════════════════════════════════════════════════
	// AFTER MODIFICATION
	// ════════════════════════════════════════════════
	fmt.Println("After slice change:")
	fmt.Println("menu:", menu)
	// Output: menu: [Matcha Coffee Juice]
	//                ^^^^^^ Changed!
	fmt.Println("slice:", slice)
	// Output: slice: [Matcha Coffee]
	//                 ^^^^^^ Changed!

	// ════════════════════════════════════════════════
	// KEY POINT: Slice is a REFERENCE, not a COPY
	// ════════════════════════════════════════════════
	// Changing slice[0] also changes menu[0]
	// They point to the same underlying data
}


/*

## 🔑 Important Concept - Slice References Array:
```
Original Array:
menu = [Tea, Coffee, Juice]
        ↑     ↑
        │     │
slice = [Tea, Coffee]  ← Points to same memory

After slice[0] = "Matcha":
menu = [Matcha, Coffee, Juice]
        ↑       ↑
        │       │
slice = [Matcha, Coffee]  ← Same memory changed!

*/
