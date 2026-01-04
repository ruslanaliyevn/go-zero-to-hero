package main

import "fmt"

func main() {
	// ════════════════════════════════════════════════
	// ARRAY LITERAL - Initialize with values
	// ════════════════════════════════════════════════
	// Declare and initialize array in one line
	coffeeTypes := [3]string{"Espresso", "Latte", "Cappuccino"}
	fmt.Println("Types of coffee:", coffeeTypes)
	// Output: Types of coffee: [Espresso Latte Cappuccino]

	// ════════════════════════════════════════════════
	// ARRAY LENGTH
	// ════════════════════════════════════════════════
	// len() returns the number of elements in array
	fmt.Println("Length of the array:", len(coffeeTypes))
	// Output: Length of the array: 3

	// ════════════════════════════════════════════════
	// ACCESSING LAST ELEMENT
	// ════════════════════════════════════════════════
	// Use len()-1 to access last element
	// Array indices: 0, 1, 2 (length is 3)
	// Last index: len(coffeeTypes)-1 = 3-1 = 2
	coffeeTypes[len(coffeeTypes)-1] = "Milk"
	fmt.Println("Types of coffee:", coffeeTypes)
	// Output: Types of coffee: [Espresso Latte Milk]

	// ════════════════════════════════════════════════
	// len() with STRINGS
	// ════════════════════════════════════════════════
	// len() also works with strings - returns character count
	fmt.Println("String length is:", len("This is a coffee string!"))
	// Output: String length is: 24
}
