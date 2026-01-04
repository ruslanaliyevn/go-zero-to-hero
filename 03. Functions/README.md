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

**Important:** 
- Parameters only exist inside the function
- Order, type, and count must match when calling

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

**Why use return?**
- Need the result for further calculations
- Want to use the value elsewhere in your code
- Function produces a value you need to store

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

**Benefits:**
- Variables are automatically created
- Self-documenting code
- Can use "naked return" (just `return`)

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

**Use when:**
- Need a temporary function
- Function is used only once or in specific scope

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

**How it works:**
- Inner function "captures" outer variables
- Variables survive after outer function ends
- Each call remembers the previous state

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

## Key Takeaways

✅ **Functions** organize code into reusable blocks  
✅ **Parameters** are function inputs (local scope only)  
✅ **Return** sends values back to the caller  
✅ **Multiple returns** useful for related data (e.g., result + error)  
✅ **Named returns** make code self-documenting  
✅ **Anonymous functions** for temporary use  
✅ **Closures** capture and remember outer variables between calls

---

## Next Steps

After learning functions, explore:
- Pointers (pass by reference vs pass by value)
- Structs (custom data types)
- Methods (functions attached to types)

---

## License

Personal learning repository.
