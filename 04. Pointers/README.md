# Go Learning Notes

---

## Pointers

### What is a Pointer?

A **pointer** stores the **memory address** of a variable, not the value itself.

---

### `&` Operator - Get Address
```go
coffee := "Espresso"
pointer := &coffee

fmt.Println(coffee)   // "Espresso"
fmt.Println(pointer)  // 0xc000010230
```

---

### Copy vs Pointer
```go
// Copy - different addresses
coffee := "Espresso"
copy := coffee

fmt.Printf("%p\n", &coffee)  // 0xc000010230
fmt.Printf("%p\n", &copy)    // 0xc000010250 (DIFFERENT)

// Pointer - same address
pointer := &coffee
fmt.Printf("%p\n", pointer)  // 0xc000010230 (SAME)
```

---

### `*` Operator - Dereference (Change Value)
```go
price := 4.50

// Create pointer
var ptr *float64 = &price

// Change value through pointer
*ptr = 7.50

fmt.Println(price)  // 7.50 (changed!)
```

---

### How It Works
```
Memory: 0xc00000a088
Value:  4.50 → 7.50

price  ← original variable
*ptr   ← changes through pointer (same memory)
```

---

### Key Points

| Operator | Meaning |
|----------|---------|
| `&` | Get address |
| `*` (in type) | Pointer type |
| `*` (in code) | Access/change value |

✅ `&variable` → get address  
✅ `*pointer` → change value at address  
✅ Same memory = same variable

---
