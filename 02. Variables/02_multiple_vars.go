package main

import "fmt"

func main() {
	// Declare and initialize wiht var with exceplicit types
	var coffeeName string = "Espresso"

	// Type inferred
	var size = "Small"

	// Short declaration and initialization. Possible only inside functions
	price := 2.50

	fmt.Println("Medium Espresso price is $3.0")
	fmt.Println(size, coffeeName, "price is $", price)
	fmt.Printf("%s %s price is $%.2f\n", size, coffeeName, price)
	fmt.Println("Done")
}
