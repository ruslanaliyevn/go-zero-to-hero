package main

import "fmt"

func main() {
	// ============================================
	// CONST (Constant) - Immutable Values
	// ============================================
	// Store name - defined once, cannot be changed
	const shopName = "Brew & Beans"
	
	// Other const examples:
	const maxCustomers = 50        // Maximum customer capacity
	const taxRate = 0.18          // VAT rate (18%)
	const ownerName = "Ruslan"    // Owner's name
	
	// ❌ ERROR: reassigning a constant is forbidden
	// shopName = "Latte Palace"  
	// Reason: const is immutable (cannot be changed)
	
	// ============================================
	// VAR (Variable) - Mutable Values
	// ============================================
	var currentCustomers = 0       // Current customers in the store
	currentCustomers = 25          // ✅ This works because it's a var
	
	// ============================================
	// OUTPUT
	// ============================================
	fmt.Println("Welcome to", shopName)
	fmt.Println("Owner:", ownerName)
	fmt.Println("Current customers:", currentCustomers, "/", maxCustomers)
}
