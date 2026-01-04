# Pointers in Go

## What is a Pointer?

A **pointer** is a variable that stores the **memory address** of another variable, not the value itself.

---

## Key Operator: `&` (Address-of)

| Operator | Name | Description | Example |
|----------|------|-------------|---------|
| `&` | Address-of | Gets the memory address of a variable | `&coffee` |

---

## Basic Usage
```go
// Create a variable
coffee := "Espresso"

// Get its memory address using &
pointer := &coffee

// Display value
fmt.Println(coffee)    // "Espresso" (the value)

// Display address
fmt.Println(pointer)         // 0xc000010230 (memory address)
fmt.Printf("%p\n", pointer)  // 0xc000010230 (formatted address)
```

---

## Copy vs Pointer

### Copy - Creates NEW memory location
```go
coffee := "Espresso"
coffeeCopy := coffee  // Creates a copy at different address

fmt.Printf("%p\n", &coffee)      // 0xc000010230
fmt.Printf("%p\n", &coffeeCopy)  // 0xc000010250 (DIFFERENT!)
```

### Pointer - References SAME memory location
```go
coffee := "Espresso"
pointer := &coffee  // Points to coffee's address

fmt.Printf("%p\n", &coffee)  // 0xc000010230
fmt.Printf("%p\n", pointer)  // 0xc000010230 (SAME!)
```

---

## Memory Visualization
```
Memory Address    Value         Variable
──────────────    ─────────     ────────
0xc000010230  →  "Espresso"    coffee
0xc000010240  →  0xc000010230  pointer (stores address)
0xc000010250  →  "Espresso"    coffeeCopy (different location)
```

---

## Key Takeaways

✅ Every variable has a memory address  
✅ `&variable` gives you the address  
✅ **Copy** = new variable, new address, independent  
✅ **Pointer** = stores address, references same location
