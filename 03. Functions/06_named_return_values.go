package main

import "fmt"

// ============================================
// FUNCTION WITH NAMED RETURN VALUES
// ============================================
// estimateBrewTime calculates total brewing time for coffee
// Input: cupQty (int) - number of cups to brew
//        secondsPerCup (int) - time needed per cup
// Output: Named returns - already declared as variables
//         totalTimeSeconds (int) - total time in seconds
//         info (string) - description message
func estimateBrewTime(cupQty int, secondsPerCup int) (totalTimeSeconds int, info string) {
	//                                                 └──────────────┬──────────────┘
	//                                        NAMED RETURN VALUES
	//                                        These variables are automatically created by Go
	//                                        No need to declare them with var or :=
	
	// ============================================
	// ASSIGN VALUES TO NAMED RETURN VARIABLES
	// ============================================
	// Calculate total time: cups × seconds per cup
	totalTimeSeconds = cupQty * secondsPerCup
	// Example: 12 cups × 20 seconds = 240 seconds
	
	// Set description message
	info = "Estimated total brew time"
	
	// ============================================
	// NAKED RETURN
	// ============================================
	// Just write "return" with nothing after it
	// Go automatically returns all named return variables
	// Equivalent to: return totalTimeSeconds, info
	return
}

// ============================================
// MAIN FUNCTION
// ============================================
func main() {
	// Calculate brew time for 12 cups at 20 seconds per cup
	// 12 × 20 = 240 seconds
	brewTime, info := estimateBrewTime(12, 20)
	//                └───────────┬───────────┘
	//                   Function returns (240, "Estimated total brew time")
	
	// Print the results
	fmt.Println(info, brewTime)
	// Output: Estimated total brew time 240
}
