package main

import "fmt"

// ============================================
// CREATE TEMPERATURE ADJUSTER FUNCTION
// ============================================
// Returns two values:
// 1. A function that can adjust temperature (closure)
// 2. The original base temperature value (copy)
func createTemperatureAdjuster() (func(change float64) float64, float64) {
	
	// ============================================
	// LOCAL VARIABLE - "Captured" by closure
	// ============================================
	// This variable will be "remembered" by the returned function
	// Even after createTemperatureAdjuster() finishes execution
	baseTemperature := 90.0
	
	// ============================================
	// ANONYMOUS FUNCTION WITH CLOSURE
	// ============================================
	// This function "captures" baseTemperature from outer scope
	// Each call modifies the SAME baseTemperature variable
	adjustTemperature := func(change float64) float64 {
		// ════════════════════════════════════════
		// CLOSURE IN ACTION
		// ════════════════════════════════════════
		// Modifies the captured baseTemperature
		// This value persists between function calls!
		baseTemperature = baseTemperature + change
		//                └───────┬───────┘
		//          This is the CAPTURED variable
		//          NOT a copy - the actual variable!
		
		return baseTemperature
	}
	
	// ============================================
	// RETURN FUNCTION AND ORIGINAL VALUE
	// ============================================
	// Return 1: adjustTemperature function (has closure over baseTemperature)
	// Return 2: Current value of baseTemperature (COPY, not reference)
	return adjustTemperature, baseTemperature
	//     └───────┬───────┘  └──────┬──────┘
	//             │                  └─ Returns 90.0 (copy)
	//             └─ Returns function that "remembers" baseTemperature
}

// ============================================
// MAIN FUNCTION
// ============================================
func main() {
	// ════════════════════════════════════════════════
	// GET ADJUSTER FUNCTION AND ORIGINAL TEMP
	// ════════════════════════════════════════════════
	adjustTemp, OriginalTemp := createTemperatureAdjuster()
	//          └──────┬──────┘
	//                 └─ This is 90.0 (COPY, won't change)
	
	fmt.Printf("Original temperature is %.1f\n", OriginalTemp)
	// Output: Original temperature is 90.0
	
	// ════════════════════════════════════════════════
	// CALL 1: Add 1.5 degrees
	// ════════════════════════════════════════════════
	// baseTemperature: 90.0 + 1.5 = 91.5
	fmt.Printf("Adjusted Temp +1.5: %.1f grad C\n", adjustTemp(1.5))
	// Output: Adjusted Temp +1.5: 91.5 grad C
	
	// ════════════════════════════════════════════════
	// CALL 2: Subtract 3.0 degrees
	// ════════════════════════════════════════════════
	// baseTemperature: 91.5 + (-3.0) = 88.5
	// ↑ Notice: Uses PREVIOUS value (91.5), not original (90.0)
	// This proves closure is working - function "remembers" state!
	fmt.Printf("Adjusted Temp -3.0: %.1f grad C\n", adjustTemp(-3.0))
	// Output: Adjusted Temp -3.0: 88.5 grad C
	
	// ════════════════════════════════════════════════
	// CALL 3: Add 5.0 degrees
	// ════════════════════════════════════════════════
	// baseTemperature: 88.5 + 5.0 = 93.5
	// ↑ Again uses PREVIOUS value (88.5)
	fmt.Printf("Adjusted Temp +5.0: %.1f grad C\n", adjustTemp(5.0))
	// Output: Adjusted Temp +5.0: 93.5 grad C
	
	// ════════════════════════════════════════════════
	// CHECK ORIGINAL TEMPERATURE
	// ════════════════════════════════════════════════
	// OriginalTemp is still 90.0 because it's a COPY
	// The closure modifies baseTemperature inside the function,
	// not the OriginalTemp variable in main()
	fmt.Printf("Original temperature is %.1f\n", OriginalTemp)
	// Output: Original temperature is 90.0 (unchanged!)
}
