# 02. Variables in Go 📦

Understanding how to store and manage data in Go programs.

---

## 📝 The Problem - Code Repetition

Let's start with a common mistake:

```go
package main

import "fmt"

func main() {
    fmt.Println("Medium Espresso price is 2.50")
    fmt.Println("Medium Espresso price is 2.50")
    fmt.Println("Medium Espresso price is 2.50")
}
```

**Output:**
```
Medium Espresso price is 2.50
Medium Espresso price is 2.50
Medium Espresso price is 2.50
```

**Problem:** What if the price changes? You'd have to update it in three places! That's inefficient and error-prone.

---

## ✅ The Solution - Using Variables

```go
package main

import "fmt"

func main() {
    var info = "Medium Espresso price is 2.50"
    fmt.Println(info)
    fmt.Println(info)
    fmt.Println(info)
}
```

**Output:**
```
Medium Espresso price is 2.50
Medium Espresso price is 2.50
Medium Espresso price is 2.50
```

**Advantage:** Change the value once, and it updates everywhere! This is the power of variables.

---

## 🔄 Variable Reassignment

Variables can change their values during program execution:

```go
package main

import "fmt"

func main() {
    var info = "Medium Espresso price is 2.50"
    fmt.Println(info)
    
    info = "Large Espresso price is $4.0"
    fmt.Println(info)
    fmt.Println(info)
}
```

**Output:**
```
Medium Espresso price is 2.50
Large Espresso price is $4.0
Large Espresso price is $4.0
```

**Key Point:** You can reassign variables to new values of the **same type**.

---

## 🚫 Type Safety in Go

Go is **statically typed** - once a variable's type is determined, it stays that way forever!

### What This Means

```go
var info = "text"       // ✅ Go infers: info is a string
info = "new text"       // ✅ Correct - still a string
info = 10               // ❌ ERROR! Cannot change string to int
```

### Why Does This Happen?

When you write `var info = "text"`:
1. Go sees the value is text (in quotes)
2. Go automatically decides: "This is a **string** type"
3. From now on, `info` can ONLY hold strings
4. You cannot change its type to int, float, or anything else

### More Examples

```go
var age = 25           // Go infers: int
age = 30               // ✅ Works - both are integers
age = "thirty"         // ❌ ERROR! age is int, not string

var price = 9.99       // Go infers: float64
price = 12.50          // ✅ Works - both are floats
price = 10             // ⚠️ Be careful! This might work but can cause issues
```

**Remember:** Type is decided at declaration and never changes!

---

## 🗂️ Working with Multiple Files

### The Problem - Duplicate `main()` Functions

Imagine you create two Go files in the same folder:

**File 1: `01_basic_declaration.go`**
```go
package main
import "fmt"
func main() {
    fmt.Println("First program")
}
```

**File 2: `02_multiple_vars.go`**
```go
package main
import "fmt"
func main() {
    fmt.Println("Second program")
}
```

**When you try to run:**
```bash
go run *.go
```

**Error:** `other declaration of main`

### Why Does This Happen?

- Each Go **package** can have only ONE `main()` function
- When files are in the same directory, they belong to the same package
- Two `main()` functions = Conflict! Go doesn't know which one to run

### ✅ Solution: Separate Directories

Create a folder for each program:

```
02. Variables/
├── 01_basic_declaration/
│   └── 01_basic_declaration.go    ← Has its own main()
└── 02_multiple_vars/
    └── 02_multiple_vars.go        ← Has its own main()
```

Now they are **separate packages** - no more conflicts!

**Run them individually:**
```bash
# Run first program
go run 01_basic_declaration/01_basic_declaration.go

# Run second program  
go run 02_multiple_vars/02_multiple_vars.go
```

**Key Rule:** One directory = One package = One `main()` function

---

## 🎯 Three Ways to Declare Variables

Go gives you three options for declaring variables, each with its own use case:

```go
package main

import "fmt"

func main() {
    // Method 1: Explicit type declaration
    var coffeeName string = "Espresso"
    
    // Method 2: Type inference with var
    var size = "Small"
    
    // Method 3: Short declaration (only inside functions!)
    price := 2.50
    
    fmt.Println(size, coffeeName, "price is $", price)
}
```

**Output:**
```
Small Espresso price is $ 2.50
```

### Declaration Methods Comparison

| Method | Syntax | Type Declaration | Where Can Use | Best For |
|--------|--------|------------------|---------------|----------|
| **Explicit var** | `var name type = value` | Required | Anywhere | When type clarity is important |
| **Inferred var** | `var name = value` | Auto-detected | Anywhere | Package-level variables |
| **Short declaration** | `name := value` | Auto-detected | **Functions only** | Quick, clean code inside functions |

### 🔑 Critical Rule

The `:=` short declaration can **ONLY** be used inside functions!

```go
// ❌ This doesn't work at package level
package main
price := 2.50  // ERROR!

// ✅ This works
package main
var price = 2.50  // Correct!

func main() {
    discount := 0.10  // ✅ Short declaration works here!
}
```

---

## 🎨 String Formatting

### Basic Printing

```go
package main

import "fmt"

func main() {
    var coffeeName string = "Espresso"
    var size = "Small"
    price := 2.50
    
    fmt.Println(size, coffeeName, "price is $", price)
}
```

**Output:**
```
Small Espresso price is $ 2.5
```

**Problem:** Notice the spacing? `Println` adds spaces between arguments, and the price shows as `2.5` instead of `2.50`.

### Formatted Printing (Better!)

```go
package main

import "fmt"

func main() {
    var coffeeName string = "Espresso"
    var size = "Small"
    price := 2.50
    
    // Basic printing
    fmt.Println(size, coffeeName, "price is $", price)
    
    // Formatted printing
    fmt.Printf("%s %s price is $%.2f\n", size, coffeeName, price)
    
    fmt.Println("Done")
}
```

**Output:**
```
Small Espresso price is $ 2.5
Small Espresso price is $2.50
Done
```

**Much better!** The second line is properly formatted with exactly 2 decimal places.

---

## 📐 Printf Format Specifiers

When using `Printf`, you use special codes to format your output:

| Specifier | Type | Description | Example Output |
|-----------|------|-------------|----------------|
| `%s` | string | Text/string value | `Espresso` |
| `%d` | int | Whole number | `42` |
| `%f` | float | Decimal number | `2.500000` |
| `%.2f` | float | Decimal with 2 places | `2.50` |
| `%v` | any | Default format (works for any type) | Varies |
| `%T` | any | Shows the type | `string`, `int`, etc. |
| `\n` | - | New line | Line break |

### Printf vs Println

**Println:**
- ✅ Automatically adds spaces between arguments
- ✅ Automatically adds newline at the end
- ❌ No control over formatting

**Printf:**
- ✅ Complete control over formatting
- ✅ Precise number formatting
- ❌ Must manually add `\n` for newlines

---

## 💡 Important Concepts

### 1. Zero Values

Variables declared without initialization get **default values**:

```go
var name string    // "" (empty string)
var age int        // 0
var price float64  // 0.0
var isActive bool  // false
```

### 2. Unused Variables = Error

Go is strict about unused variables:

```go
func main() {
    var unused = "test"  // ❌ Error: declared but not used
}
```

**Solution:** Either use it or remove it!

### 3. Multiple Declarations

You can declare multiple variables at once:

```go
// Same type
var x, y, z int = 1, 2, 3

// Different types
var (
    name    string  = "Coffee"
    price   float64 = 2.50
    inStock bool    = true
)

// Short declaration (inside functions)
a, b, c := 1, "two", 3.0
```

### 4. Naming Conventions

**Good variable names:**
- ✅ `coffeeName` (camelCase)
- ✅ `userAge`
- ✅ `totalPrice`

**Exported variables (accessible from other packages):**
- ✅ `MaxConnections` (starts with uppercase)
- ✅ `DefaultTimeout`

**Keep names:**
- Descriptive but not too long
- Meaningful to anyone reading the code
- Consistent throughout your project

---

## 🎯 Best Practices

1. **Use `:=` inside functions** - It's cleaner and idiomatic Go
2. **Use `var` at package level** - No choice here, `:=` doesn't work there
3. **Use explicit types when clarity matters** - Sometimes being explicit is better
4. **Use `Printf` for formatted output** - Especially with numbers and currency
5. **One `main()` per directory** - Organize files properly
6. **Always use declared variables** - Go won't compile otherwise
7. **Keep names meaningful** - Your future self will thank you

---

## 🧠 What I Learned

### Variable Basics
- ✅ Variables store data for reuse
- ✅ Variables can be reassigned (same type only)
- ✅ Go automatically infers types when not specified
- ✅ Unused variables cause compilation errors

### Declaration Methods
- ✅ `var name type = value` - Explicit and clear
- ✅ `var name = value` - Let Go figure out the type
- ✅ `name := value` - Quick shorthand (functions only!)

### Type System
- ✅ Go is statically typed (types can't change)
- ✅ Type inference saves typing but keeps safety
- ✅ Zero values prevent uninitialized variables

### Formatting Output
- ✅ `Println()` for simple output
- ✅ `Printf()` for controlled formatting
- ✅ Format specifiers give precise control

### Project Organization
- ✅ One `main()` function per package
- ✅ Separate directories for separate programs
- ✅ Proper structure prevents conflicts

---

## 🎯 Summary

### What Are Variables?
Variables are containers that store information in your program. Instead of writing the same value multiple times, you store it once and reuse it everywhere.

### Why Use Variables?
✅ **Reusability** - Write once, use many times  
✅ **Flexibility** - Change the value in one place, updates everywhere  
✅ **Readability** - `price` is clearer than `2.50` scattered in code  
✅ **Maintainability** - Easier to update and debug your code

### Three Ways to Declare
```go
var name string = "Coffee"  // 1. Explicit type (clear but verbose)
var name = "Coffee"         // 2. Type inference (Go figures it out)
name := "Coffee"            // 3. Short form (only in functions!)
```

### Golden Rules
1. **Type is fixed** - Once a variable is a string, it stays a string
2. **Use it or lose it** - Unused variables cause errors
3. **`:=` only in functions** - Outside functions, use `var`
4. **One `main()` per package** - Separate programs into different folders

### Format Your Output
- `Println()` for quick, simple printing
- `Printf()` for precise control and formatting

---

*Variables transform hardcoded chaos into organized, maintainable code!* 📦✨
