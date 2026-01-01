package main

import "fmt"

func main() {
	// ============================================
	// UNTYPED CONSTANT - Type Adaptation
	// ============================================
	// Untyped constant with integer value
	// Go compiler will automatically adapt this to the required type
	const rewardPoints = 10
	
	// Print the default type of the constant
	// Output: int (Go's default type for integer constants)
	fmt.Printf("Default type of rewardPoints is %T\n", rewardPoints)

	// ============================================
	// TYPED VARIABLE - float64
	// ============================================
	// Variable to store total reward points (with decimal values)
	var totalRewardPoints float64 = 150.3

	// ============================================
	// TYPE ADAPTATION IN ACTION
	// ============================================
	// Adding untyped constant to a float64 - VALID ✅
	// The constant 'rewardPoints' automatically adapts from int to float64
	// This is the beauty of untyped constants in Go!
	totalRewardPoints = totalRewardPoints + rewardPoints

	// Print the updated total with 2 decimal places
	// %.2f formats the float to 2 decimal places
	fmt.Printf("Updated loyalty points: %.2f\n", totalRewardPoints)
}
