package main 

import "fmt"

// ============================================
// UPDATE TOTAL POINTS
// ============================================
// Adds new points to existing points
// Input: currentPoints (int) - current total points
//        newPoints (int) - newly earned points
// Output: int - updated total points
func updateTotalPoints(currentPoints int, newPoints int) int {
    // Simply add the two values and return the sum
    return currentPoints + newPoints
}

// ============================================
// CALCULATE LOYALTY POINTS
// ============================================
// Calculates loyalty points from amount spent
// Formula: spending × 2 = points
// Input: amountSpent (float64) - money spent
// Output: int - loyalty points earned
func calculateLoyaltyPoints(amountSpent float64) int {
    // Convert float64 result to int (truncates decimals)
    loyaltyPoints := int(amountSpent * 2)
    return loyaltyPoints
}

// ============================================
// MAIN FUNCTION
// ============================================
func main() {
    // ============================================
    // STEP 1: Initialize starting points
    // ============================================
    totalPoints := 120  // Customer starts with 120 points
    
    // ============================================
    // STEP 2: Calculate new points from spending
    // ============================================
    // Customer spent $9.50 → earns 19 points
    var newlyEarnedPoints int = calculateLoyaltyPoints(9.50)
    fmt.Println("Earned points today:", newlyEarnedPoints)
    // Output: Earned points today: 19
    
    // ============================================
    // STEP 3: Update total points
    // ============================================
    // Add newly earned points to existing total
    // 120 (old total) + 19 (new points) = 139 (new total)
    totalPoints = updateTotalPoints(totalPoints, newlyEarnedPoints)
    //            └──────────┬─────────┘
    //                       └─── Function returns 139, stored back in totalPoints
    
    fmt.Println("Updated loyalty points:", totalPoints)
    // Output: Updated loyalty points: 139
}


// ## 🔍 Execution Flow:
// ```
// ┌─────────────────────────────────────────────────┐
// │ STEP-BY-STEP EXECUTION                          │
// ├─────────────────────────────────────────────────┤
// │                                                 │
// │ 1️⃣ totalPoints = 120                            │
// │    Customer starts with 120 points              │
// │                                                 │
// │ 2️⃣ calculateLoyaltyPoints(9.50)                 │
// │    9.50 × 2 = 19.0 → int(19.0) = 19            │
// │    newlyEarnedPoints = 19                       │
// │                                                 │
// │ 3️⃣ updateTotalPoints(120, 19)                   │
// │    120 + 19 = 139                               │
// │    totalPoints = 139                            │
// │                                                 │
// │ 4️⃣ Print: "Updated loyalty points: 139"        │
// └─────────────────────────────────────────────────┘
// 
