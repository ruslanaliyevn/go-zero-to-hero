package main

import "fmt"

// ════════════════════════════════════════════════
// SWITCH STATEMENT - Multi-Way Selection
// ════════════════════════════════════════════════
// Cleaner alternative to multiple if-else if chains
// Compares a value against multiple cases
// Syntax:
//   switch variable {
//   case value1:
//     code
//   case value2, value3:
//     code (multiple values in one case)
//   default:
//     code (if no match)
//   }

func main() {
	// ════════════════════════════════════════════════
	// DAILY PROMOTION SYSTEM
	// ════════════════════════════════════════════════
	day := "Sunday"

	// ════════════════════════════════════════════════
	// SWITCH - Compare 'day' against each case
	// ════════════════════════════════════════════════
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Friday":
		// Multiple values in one case (separated by commas)
		// Matches if day equals ANY of these values
		// "Sunday" ≠ any of these → skip
		fmt.Println("Weekday special: Buy one get one 50% off")
		
	case "Saturday", "Sunday":
		// "Sunday" == "Sunday" → MATCH! ✓
		fmt.Println("Weekend special: Free croissant with any coffee!")
		// ⚠️ Automatic break - doesn't fall through to next case
		
	default:
		// Executes if NO cases match
		// Similar to 'else' in if-else
		fmt.Println("Unknown day")
	}
	// Output: "Weekend special: Free croissant with any coffee!"

	// ════════════════════════════════════════════════
	// KEY DIFFERENCES: SWITCH vs IF-ELSE
	// ════════════════════════════════════════════════
	// ✓ No 'break' needed (automatic in Go)
	// ✓ Multiple values per case with comma
	// ✓ More readable for many conditions
	// ✓ Can only compare equality (not >, <, etc.)
	// ✓ 'default' case is optional
	
	// ════════════════════════════════════════════════
	// EQUIVALENT IF-ELSE VERSION (more verbose)
	// ════════════════════════════════════════════════
	// if day == "Monday" || day == "Tuesday" || day == "Wednesday" || day == "Friday" {
	//     fmt.Println("Weekday special...")
	// } else if day == "Saturday" || day == "Sunday" {
	//     fmt.Println("Weekend special...")
	// } else {
	//     fmt.Println("Unknown day")
	// }
}
