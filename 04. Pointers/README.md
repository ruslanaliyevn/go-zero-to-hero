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

### Functions: With vs Without Pointers

#### WITHOUT Pointer - Returns New Value
```go
func calculateDiscount(price float64, rate float64) float64 {
    return price - (price * rate)
}

// Usage:
coffeePrice := 5.00
coffeePrice = calculateDiscount(coffeePrice, 0.10)  // Must capture return
fmt.Println(coffeePrice)  // 4.50
```

**How it works:**
- Function receives a **copy** of the value
- Original variable **unchanged**
- Must **return** new value
- Must **capture** the returned value

#### WITH Pointer - Modifies Original
```go
func applyDiscount(price *float64, rate float64) {
    *price = *price - (*price * rate)
}

// Usage:
coffeePrice := 5.00
applyDiscount(&coffeePrice, 0.10)  // Original modified directly
fmt.Println(coffeePrice)  // 4.50
```

**How it works:**
- Function receives **memory address**
- Original variable **modified directly**
- No return needed
- No need to capture result

---

### Comparison

| Feature | WITHOUT Pointer | WITH Pointer |
|---------|----------------|--------------|
| Parameter | Value (copy) | Address (`*float64`) |
| Original variable | Unchanged | Changed |
| Return value | Required | Not needed |
| Usage | `func(value)` | `func(&value)` |
| When to use | Need new value | Modify original |

---

### Key Points

✅ `&variable` → get address  
✅ `*pointer` → access/modify value  
✅ **Without pointer**: copy, return, capture  
✅ **With pointer**: address, modify, no return

---

## License

Personal learning repository.
