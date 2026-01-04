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

---

### DevOps Use Case: Server Configuration Management

**Why Pointers Matter in DevOps:**

In real DevOps scenarios, you often need to modify configurations directly without creating copies. This is more efficient for:
- Managing multiple servers
- Updating configurations in real-time
- Reducing memory usage
- Improving performance

**Example: Server Scaling**
```go
type ServerConfig struct {
    Name     string
    CPUCores int
    MemoryGB int
    DiskGB   int
}

// Upgrade server using pointer - modifies original
func upgradeServer(config *ServerConfig, newCPU int, newRAM int) {
    config.CPUCores = newCPU
    config.MemoryGB = newRAM
}

// Usage:
prodServer := ServerConfig{
    Name:     "production-api",
    CPUCores: 4,
    MemoryGB: 8,
    DiskGB:   200,
}

// Scale up when traffic increases
upgradeServer(&prodServer, 8, 16)   // Directly modifies prodServer
upgradeServer(&prodServer, 16, 32)  // Scale again if needed
```

**Benefits:**
- ✅ No need to return and capture values
- ✅ Direct modification of server config
- ✅ Efficient for managing multiple servers
- ✅ Cleaner code in production scenarios

---

### Key Points

✅ `&variable` → get address  
✅ `*pointer` → access/modify value  
✅ **Without pointer**: copy, return, capture  
✅ **With pointer**: address, modify, no return  
✅ **DevOps**: pointers for efficient config management

---

## License

Personal learning repository.
