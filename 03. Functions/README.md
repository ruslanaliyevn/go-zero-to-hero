# Go Functions

Learn functions in Go with practical examples from a coffee shop management system.

---

## Table of Contents

1. [Basic Function](#basic-function)
2. [Parameters](#parameters)
3. [Return Values](#return-values)
4. [Multiple Returns](#multiple-returns)
5. [Named Returns](#named-returns)
6. [Anonymous Functions](#anonymous-functions)
7. [Closures](#closures)
8. [Best Practices](#best-practices)

---

## Basic Function

A function with no inputs and no outputs - just performs an action.
```go
func greet() {
    fmt.Println("Welcome to Brew & Beans")
}

// Usage
greet()  // Prints the welcome message
greet()  // Can call multiple times
```

**When to use:**
- Perform an action without needing data in or out
- Display messages, log information
- Execute simple operations

---

## Parameters

Functions can accept input values (parameters).

### Single Parameter
```go
func greet(shopName string) {
    fmt.Println("Welcome to", shopName)
}

greet("Brew & Beans")  // Pass value as argument
```

### Multiple Parameters
```go
func getDrinkInfo(name string, drink string, price float64) {
    fmt.Printf("%s's drink: %s at $%.2f\n", name, drink, price)
}

getDrinkInfo("Alice", "Latte", 4.50)
```

### Parameter Scope
```go
func greet(shopName string) {
    var message = "Welcome to"  // Local variable
    fmt.Println(message, shopName)
}

// ❌ shopName and message don't exist here
// They only exist inside the greet() function
```

**Important:** 
- Parameters only exist inside the function (local scope)
- Order, type, and count must match when calling
- Parameters are copies of the values passed in

---

## Return Values

Functions can send results back to the caller.
```go
func calculatePoints(spent float64) int {
    return int(spent * 2)
}

// Capture the returned value
points := calculatePoints(9.50)  // points = 19
fmt.Println(points)
```

### Why use return?

- ✅ Need the result for further calculations
- ✅ Want to use the value elsewhere in your code
- ✅ Function produces a value you need to store
- ✅ Makes functions more flexible and reusable

### Return vs No Return
```go
// WITHOUT return - just prints
func displayTotal(amount float64) {
    fmt.Printf("Total: $%.2f\n", amount)
}

// WITH return - returns value for use
func calculateTotal(amount float64) float64 {
    return amount * 1.10  // Add 10% tax
}

// Usage difference
displayTotal(100.0)           // Just prints, no value to capture
total := calculateTotal(100.0) // Returns 110.0, can use it
```

---

## Multiple Returns

Go functions can return multiple values at once.
```go
func processPayment(total float64, paid float64) (float64, float64) {
    due := total
    change := paid - total
    return due, change  // Return two values
}

// Capture both values
totalDue, change := processPayment(8.50, 10.00)
fmt.Println(totalDue)  // 8.50
fmt.Println(change)    // 1.50

// Ignore unwanted values with _
totalDue, _ := processPayment(8.50, 10.00)  // Ignore change
```

### Common Pattern: Value + Error
```go
func divide(a, b float64) (float64, string) {
    if b == 0 {
        return 0, "Error: division by zero"
    }
    return a / b, ""  // Empty string = no error
}

// Usage
result, err := divide(10, 2)
if err != "" {
    fmt.Println(err)
} else {
    fmt.Printf("Result: %.2f\n", result)
}
```

---

## Named Returns

You can name return values in the function signature.
```go
func estimateTime(cups int, secPerCup int) (totalSeconds int, info string) {
    totalSeconds = cups * secPerCup  // No need for var or :=
    info = "Estimated brew time"
    return  // "Naked return" - automatically returns named values
}

time, msg := estimateTime(12, 20)
fmt.Println(msg, time)  // Estimated brew time 240
```

### Benefits

- ✅ Variables are automatically created
- ✅ Self-documenting code (clear what function returns)
- ✅ Can use "naked return" (just `return`)
- ✅ Reduces repetition in return statements

### When to use
```go
// ✅ Good for small functions with clear purpose
func calculateStats(price float64) (subtotal float64, tax float64, total float64) {
    subtotal = price
    tax = price * 0.10
    total = subtotal + tax
    return
}

// ⚠️ Avoid in large functions - can be confusing
// Better to use explicit return in complex functions
```

---

## Anonymous Functions

Functions without names, stored in variables.
```go
// Create anonymous function
calculateTax := func(amount float64) float64 {
    return amount * 0.10
}

// Use it like a normal function
tax := calculateTax(100.0)
fmt.Println(tax)  // 10.0
```

### Use Cases

**Temporary operations:**
```go
taxRate := 0.10

calculateTax := func(amount float64) float64 {
    return amount * taxRate  // Access outer variable
}

tax := calculateTax(100.0)
```

**Immediately invoked:**
```go
result := func(x int) int {
    return x * 2
}(5)  // Call immediately with 5

fmt.Println(result)  // 10
```

**When to use:**
- ✅ Need a temporary function
- ✅ Function is used only once or in specific scope
- ✅ Want to access outer scope variables (closure)

---

## Closures

A function that "captures" and remembers variables from its outer scope.

### Simple Example: Counter
```go
func createCounter() func() int {
    count := 0  // This variable is "captured"
    
    return func() int {
        count++      // Remembers and modifies count
        return count
    }
}

// Usage
counter := createCounter()
fmt.Println(counter())  // 1
fmt.Println(counter())  // 2
fmt.Println(counter())  // 3
// count persists between calls!
```

### How Closures Work
```
┌────────────────────────────────────┐
│ createCounter() called             │
│                                    │
│  count := 0  ← Created             │
│                                    │
│  return func() {                   │
│      count++  ← "Captures" count   │
│  }                                 │
│                                    │
│  Function returned, but count      │
│  SURVIVES because it's captured!   │
└────────────────────────────────────┘

counter()  → count = 1
counter()  → count = 2 (remembers previous!)
counter()  → count = 3 (still remembers!)
```

### Practical Example: Temperature Adjuster
```go
func createAdjuster() func(float64) float64 {
    temp := 90.0  // Captured variable
    
    return func(change float64) float64 {
        temp += change  // Modifies the captured temp
        return temp
    }
}

adjust := createAdjuster()
fmt.Println(adjust(1.5))   // 91.5
fmt.Println(adjust(-3.0))  // 88.5 (remembers previous 91.5!)
fmt.Println(adjust(5.0))   // 93.5 (remembers previous 88.5!)
```

### Multiple Independent Closures
```go
counter1 := createCounter()
counter2 := createCounter()

fmt.Println(counter1())  // 1
fmt.Println(counter1())  // 2

fmt.Println(counter2())  // 1 (independent from counter1!)
fmt.Println(counter2())  // 2
```

**Key Points:**
- Inner function "captures" outer variables
- Variables survive after outer function ends
- Each closure has its own copy of variables
- Each call remembers the previous state

---

## Best Practices

### 1. Function Naming
```go
// ✅ Good - verb + noun (what it does)
calculateTotal()
getUserName()
processPayment()

// ❌ Bad - unclear purpose
doStuff()
handle()
process()
```

### 2. Keep Functions Small
```go
// ✅ Good - one clear purpose
func calculateTax(amount float64) float64 {
    return amount * 0.10
}

// ❌ Bad - doing too much
func processEverything(data string) {
    // Parse data
    // Validate data
    // Save to database
    // Send email
    // Log results
    // ... (too much in one function!)
}
```

### 3. Use Named Returns for Clarity
```go
// ✅ Clear what's being returned
func calculatePrice(base float64) (subtotal float64, tax float64, total float64) {
    // ...
}

// vs

// Less clear
func calculatePrice(base float64) (float64, float64, float64) {
    // Which float64 is which?
}
```

### 4. Limit Anonymous Functions
```go
// ✅ Good - simple, clear use case
calculate := func(x int) int { return x * 2 }

// ❌ Bad - too complex for anonymous function
calculate := func(x, y, z int, data []string, opts map[string]interface{}) (int, error) {
    // 50 lines of code...
}
// Better to make this a named function
```

### 5. Document Function Purpose
```go
// ✅ Good - clear comment
// calculateDiscount applies a percentage discount to the given price
// and returns the discounted amount
func calculateDiscount(price float64, percent float64) float64 {
    return price * (1 - percent)
}
```

---

## Summary Table

| Feature | Syntax | When to Use |
|---------|--------|-------------|
| **Basic function** | `func greet() { }` | Perform action, no input/output |
| **With parameters** | `func greet(name string) { }` | Need input values |
| **With return** | `func calc() int { }` | Need to return a result |
| **Multiple returns** | `func pay() (float64, float64) { }` | Return related data together |
| **Named returns** | `func time() (sec int, msg string) { }` | Self-documenting code |
| **Anonymous** | `f := func() { }` | Temporary/inline functions |
| **Closure** | Function capturing outer vars | Need to remember state |

---

## License

Personal learning repository.
