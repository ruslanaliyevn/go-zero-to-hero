# Go Arrays and Slices

Understanding fixed-size arrays and dynamic slices in Go.

---

## Table of Contents

1. [Arrays](#arrays)
2. [Array Literals](#array-literals)
3. [Slices from Arrays](#slices-from-arrays)
4. [Slice References](#slice-references)
5. [Creating Slices with make()](#creating-slices-with-make)
6. [Append and Capacity](#append-and-capacity)
7. [Best Practices](#best-practices)

---

## Arrays

Fixed-size collections with elements of the same type.

### Basic Array
```go
var coffeeSizes [3]string  // Declare array with size 3
fmt.Println(coffeeSizes)   // [  ] (three empty strings)

// Assign values
coffeeSizes[0] = "Small"
coffeeSizes[1] = "Medium"
coffeeSizes[2] = "Large"

fmt.Println(coffeeSizes)   // [Small Medium Large]
```

### Key Points

- **Fixed size** declared at creation
- **Zero-based indexing** (starts at 0)
- **Default values**: `""` for strings, `0` for numbers, `false` for bools
- **Bounds checking**: Accessing `arr[5]` when size is 3 causes error
```go
// ❌ Error: index out of bounds
coffeeSizes[4] = "Extra Large"  // Array size is 3, max index is 2
```

---

## Array Literals

Initialize arrays with values directly.
```go
// Declare and initialize in one line
coffeeTypes := [3]string{"Espresso", "Latte", "Cappuccino"}
//             └┬┘ Size must match number of elements
```

### Working with Arrays
```go
// Get length
len(coffeeTypes)  // 3

// Access last element
lastIndex := len(coffeeTypes) - 1
coffeeTypes[lastIndex] = "Milk"

// len() also works with strings
len("Hello")  // 5
```

---

## Slices from Arrays

Create flexible views into arrays.

### Slice Syntax
```go
desserts := [4]string{"Muffin", "Brownie", "Croissant", "Cookie"}
//          indices:    0        1          2           3

// [start:end] - start inclusive, end exclusive
slice := desserts[1:3]   // [Brownie Croissant]
```

### Slice Variations
```go
arr := [4]string{"A", "B", "C", "D"}

arr[1:3]   // [B C]     - indices 1, 2
arr[:]     // [A B C D] - all elements
arr[2:]    // [C D]     - from index 2 to end
arr[:3]    // [A B C]   - from start to index 3 (exclusive)
```

**Pattern:**
- `[:]` → all elements
- `[n:]` → from index n to end
- `[:n]` → from start to index n (not including n)
- `[start:end]` → from start (inclusive) to end (exclusive)

---

## Slice References

**Critical:** Slices reference the underlying array - they don't copy data!
```go
menu := [3]string{"Tea", "Coffee", "Juice"}
slice := menu[:2]  // [Tea Coffee]

// Modify slice
slice[0] = "Matcha"

// Original array also changed!
fmt.Println(menu)   // [Matcha Coffee Juice]
fmt.Println(slice)  // [Matcha Coffee]
```

### How It Works
```
Array:  [Tea, Coffee, Juice]
         ↑     ↑
         │     │
Slice:  [Tea, Coffee]  ← Points to same memory

After modification:
Array:  [Matcha, Coffee, Juice]
         ↑       ↑
         │       │
Slice:  [Matcha, Coffee]  ← Same memory!
```

**Key Point:** Modifying a slice modifies the original array.

---

## Creating Slices with make()

Two ways to create slices without arrays.

### Method 1: Slice Literal
```go
// Direct initialization
ratings := []int{5, 4, 5, 5, 3}
//         └┬┘ Empty brackets = slice (not array)
```

### Method 2: make() Function
```go
// Create empty slice with specific length
coffeeTypes := make([]string, 3)
//             └────┬────┘  └┬┘
//              Function   Length

// Creates: ["", "", ""]
coffeeTypes[0] = "Cappuccino"
coffeeTypes[1] = "Latte"
coffeeTypes[2] = "Espresso"
```

### Array vs Slice
```go
// ARRAY - Fixed size
arr := [3]string{"A", "B"}  // [3] = array

// SLICE - Dynamic size  
slice := []string{"A", "B"}  // [] = slice
slice := make([]string, 3)   // Using make()
```

---

## Append and Capacity

Slices can grow dynamically using `append()`.

### Length vs Capacity
```go
slice := make([]string, 2, 5)
//                      │  └─ Capacity: 5 (allocated space)
//                      └──── Length: 2 (actual elements)

len(slice)  // 2 - number of elements
cap(slice)  // 5 - underlying array size
```

### Append Behavior
```go
menu := []string{"Cake", "Pie"}
// len=2, cap=2

menu = append(menu, "Donut")
// len=3, cap=4 (DOUBLED - new array created!)

menu = append(menu, "Ice Cream")
// len=4, cap=4 (same array, still has room)

menu = append(menu, "Cream")
// len=5, cap=8 (DOUBLED again - new array created!)
```

### Capacity Growth Pattern
```
Initial:     len=2, cap=2
Append 1:    len=3, cap=4 (doubled, new memory)
Append 2:    len=4, cap=4 (same memory)
Append 3:    len=5, cap=8 (doubled, new memory)
```

**When Go creates new array:**
- ✅ When `len == cap` and you append
- ✅ New capacity = old capacity × 2
- ✅ All elements copied to new location

### make() with Capacity
```go
// make(type, length, capacity)
cupSizes := make([]string, 0, 5)
//                         └┬┘ └┬┘
//                        len  cap

// len=0: Can't access elements yet
// cap=5: Room for 5 elements

// ❌ This fails (len is 0)
cupSizes[0] = "Small"  // ERROR!

// ✅ Use append when len is 0
cupSizes = append(cupSizes, "Small", "Medium")
// Now len=2, can access cupSizes[0] and cupSizes[1]

cupSizes[0] = "Extra small"  // ✅ OK now
```

---

## Best Practices

### 1. Choose Array or Slice
```go
// ✅ Use arrays when size is fixed and known
var weekDays [7]string

// ✅ Use slices for dynamic collections
userList := []string{}
```

### 2. Pre-allocate Capacity
```go
// ❌ Bad: Many reallocations
slice := []string{}
for i := 0; i < 1000; i++ {
    slice = append(slice, "item")
}

// ✅ Good: Pre-allocate
slice := make([]string, 0, 1000)
for i := 0; i < 1000; i++ {
    slice = append(slice, "item")  // No reallocation!
}
```

### 3. Be Aware of References
```go
// Slice references array - modification affects both
arr := [3]int{1, 2, 3}
slice := arr[:]
slice[0] = 99
// arr is now [99, 2, 3]

// To avoid: Copy the slice
arrCopy := make([]int, len(arr))
copy(arrCopy, arr[:])
```

### 4. Check Bounds
```go
slice := []int{1, 2, 3}

// ✅ Safe
if index < len(slice) {
    value := slice[index]
}

// ❌ Unsafe - can panic
value := slice[10]  // Runtime error!
```

---

## Summary

### Arrays
```go
var arr [3]string           // Fixed size
arr := [3]string{"A", "B"}  // Literal
arr[0] = "value"            // Access
len(arr)                    // Length
```

### Slices
```go
slice := []string{"A", "B"}     // Literal
slice := make([]string, 3)      // make() with length
slice := make([]string, 0, 5)   // make() with capacity
slice = append(slice, "C")      // Add element
len(slice)                      // Length
cap(slice)                      // Capacity
```

### Key Differences

| Feature | Array | Slice |
|---------|-------|-------|
| **Size** | Fixed | Dynamic |
| **Syntax** | `[3]string` | `[]string` |
| **Can grow** | ❌ No | ✅ Yes |
| **Type** | Size in type | No size in type |
| **Use when** | Size known | Size unknown/changing |

---

## Key Takeaways

✅ **Arrays** have fixed size, defined at creation  
✅ **Slices** are dynamic views into arrays  
✅ **Slice syntax**: `[start:end]` (start inclusive, end exclusive)  
✅ **Slices reference arrays** - modifications affect original  
✅ **make()** creates slices with specific length/capacity  
✅ **append()** grows slices, may reallocate memory  
✅ **Capacity doubles** when slice is full  
✅ **Pre-allocate capacity** for better performance


---

## License

Personal learning repository.
