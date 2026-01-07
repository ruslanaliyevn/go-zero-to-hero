// 07_slice_capacity_overflow.go
package main

import "fmt"

func main() {
	// ════════════════════════════════════════════════
	// ORIGINAL ARRAY
	// ════════════════════════════════════════════════
	desserts := [3]string{"Cupcake", "Eclair", "Ice cream"}
	fmt.Println("Array:", desserts)
	// Array: [Cupcake Eclair Ice cream]

	fmt.Println("-------------------")

	// ════════════════════════════════════════════════
	// CREATE SLICE FROM ARRAY
	// ════════════════════════════════════════════════
	// Slice takes first element [:1]
	// IMPORTANT: Capacity = remaining array size from start point
	slice := desserts[:1]
	// len=1 (one element)
	// cap=3 (capacity from index 0 to end of array)
	fmt.Println("Len of slice:", len(slice), "Cap of slice:", cap(slice), "Slice:", slice)
	// Len: 1, Cap: 3, Slice: [Cupcake]

	fmt.Println("-------------------")

	// ════════════════════════════════════════════════
	// FIRST APPEND - Within Capacity
	// ════════════════════════════════════════════════
	// cap=3, len=1 → has room, uses existing array
	slice = append(slice, "Macaron")
	// Overwrites desserts[1] with "Macaron"
	fmt.Println("Array:", desserts)
	// Array: [Cupcake Macaron Ice cream] ← "Eclair" replaced!
	fmt.Println("Len of slice:", len(slice), "Cap of slice:", cap(slice), "Slice:", slice)
	// Len: 2, Cap: 3, Slice: [Cupcake Macaron]

	fmt.Println("-------------------")

	// ════════════════════════════════════════════════
	// SECOND APPEND - Still Within Capacity
	// ════════════════════════════════════════════════
	// cap=3, len=2 → still has room
	slice = append(slice, "Cake")
	// Overwrites desserts[2] with "Cake"
	fmt.Println("Array:", desserts)
	// Array: [Cupcake Macaron Cake] ← "Ice cream" replaced!
	fmt.Println("Len of slice:", len(slice), "Cap of slice:", cap(slice), "Slice:", slice)
	// Len: 3, Cap: 3, Slice: [Cupcake Macaron Cake]

	fmt.Println("-------------------")

	// ════════════════════════════════════════════════
	// THIRD APPEND - CAPACITY OVERFLOW!
	// ════════════════════════════════════════════════
	// cap=3, len=3 → NO ROOM LEFT
	// Go creates NEW underlying array
	slice = append(slice, "Juice")
	// New array created: capacity doubled (3 → 6)
	// Original array NO LONGER CONNECTED to slice!
	fmt.Println("Array:", desserts)
	// Array: [Cupcake Macaron Cake] ← UNCHANGED now!
	fmt.Println("Len of slice:", len(slice), "Cap of slice:", cap(slice), "Slice:", slice)
	// Len: 4, Cap: 6, Slice: [Cupcake Macaron Cake Juice]

	fmt.Println("-------------------")

	// ════════════════════════════════════════════════
	// MODIFY SLICE AFTER OVERFLOW
	// ════════════════════════════════════════════════
	// Slice now points to NEW array
	// Modifying slice does NOT affect original array anymore
	slice[0] = "Chocolate"
	fmt.Println("Array:", desserts)
	// Array: [Cupcake Macaron Cake] ← NOT modified!
	fmt.Println("Len of slice:", len(slice), "Cap of slice:", cap(slice), "Slice:", slice)
	// Len: 4, Cap: 6, Slice: [Chocolate Macaron Cake Juice]
	//                         └────┬────┘
	//                    Only slice changed!

	fmt.Println("-------------------")
}
