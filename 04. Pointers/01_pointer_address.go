package main

import "fmt"

func main() {
	// ════════════════════════════════════════════════
	// ORIGINAL VARIABLE
	// ════════════════════════════════════════════════
	// Create a string variable
	// This is stored at some memory address (e.g., 0xc000010230)
	coffee := "Espresso"
	
	// ════════════════════════════════════════════════
	// POINTER - Get memory address
	// ════════════════════════════════════════════════
	// & operator = "address-of" operator
	// Gets the memory address where coffee is stored
	pointer := &coffee
	//         └┬┘
	//          └─ & means "give me the address of coffee"
	
	// ════════════════════════════════════════════════
	// DISPLAY ORIGINAL VARIABLE
	// ════════════════════════════════════════════════
	fmt.Println("Coffee name for coffee variable:", coffee)
	// Output: Coffee name for coffee variable: Espresso
	// Shows the VALUE stored in coffee
	
	fmt.Println("Memory location:", pointer)
	// Output: Memory location: 0xc000010230 (example address)
	// Shows the MEMORY ADDRESS (in hexadecimal format)
	
	fmt.Printf("Pointer address: %p\n", pointer)
	// Output: Pointer address: 0xc000010230
	// %p format = pointer format (shows memory address)
	// Same as above but using Printf formatting
	
	fmt.Println("-----------------------------")
	
	// ════════════════════════════════════════════════
	// COPY VARIABLE - Creates NEW memory location
	// ════════════════════════════════════════════════
	// This creates a COPY of the value
	// A NEW memory location is allocated
	coffeeCopy := coffee
	//            └──┬──┘
	//        Copies the VALUE "Espresso"
	//        Stored at DIFFERENT memory address
	
	// ════════════════════════════════════════════════
	// DISPLAY COPY VARIABLE
	// ════════════════════════════════════════════════
	fmt.Println("Coffee name for coffeeCopy variable:", coffeeCopy)
	// Output: Coffee name for coffeeCopy variable: Espresso
	// Same VALUE but different memory location
	
	fmt.Println("Memory location:", &coffeeCopy)
	//                             └────┬────┘
	//                          & gets address of coffeeCopy
	// Output: Memory location: 0xc000010250 (DIFFERENT address!)
	
	fmt.Printf("Pointer address: %p\n", &coffeeCopy)
	// Output: Pointer address: 0xc000010250
	// Shows coffeeCopy is at DIFFERENT memory location than coffee
	
	// ════════════════════════════════════════════════
	// KEY TAKEAWAY:
	// ════════════════════════════════════════════════
	// coffee      → stored at 0xc000010230
	// coffeeCopy  → stored at 0xc000010250 (DIFFERENT!)
	// 
	// They have the SAME VALUE but DIFFERENT ADDRESSES
	// This means they are INDEPENDENT copies
}
