# 01. Hello Go ☕

My very first Go program - because every programmer needs coffee!

## 📝 Code
```go
package main

import "fmt"

func main() {
	fmt.Println("I want a cup of coffee!")
}
```

## 🔍 Breaking It Down

### `package main`
Every Go program starts with a package declaration. The `main` package tells Go: "Hey, this is an executable program, not just a library!"

Think of it as the front door of your application.

### `import "fmt"`
We're bringing in the `fmt` (format) package from Go's standard library.

**fmt** gives us superpowers to:
- Print to the console
- Format strings
- Read user input

It's like Python's `print()` or C's `printf()` - but better organized.

### `func main()`
This is where the magic happens! The `main()` function is the **entry point** of every Go program.

When you run your program, Go looks for this function and starts executing from here.

**Rules:**
- Every executable Go program MUST have exactly ONE `main()` function
- It takes no parameters
- It returns nothing

### `fmt.Println()`
**Println** = Print Line

This function prints whatever you give it to the console and adds a newline at the end.

**Example:**
```go
fmt.Println("Hello")
fmt.Println("World")
```

**Output:**
```
Hello
World
```

## 🚀 How to Run

### Method 1: Direct Run
```bash
go run main.go
```
This compiles and runs your code in one step. Perfect for testing!

### Method 2: Build First
```bash
# Build executable
go build main.go

# Run it
./main       # Linux/Mac
main.exe     # Windows
```
This creates a standalone executable file you can share with anyone!

## 📤 Output
```
I want a cup of coffee!
```

## 💡 What I Learned

- ✅ Go programs need `package main` to be executable
- ✅ `import` brings in external functionality
- ✅ `func main()` is where execution begins
- ✅ `fmt.Println()` prints text to console
- ✅ Go files end with `.go`
- ✅ Go is simple and straightforward!

## 🎯 Key Takeaway

Go's structure is clean and predictable. Every program follows the same pattern:
1. Declare package
2. Import what you need
3. Write your main function
4. Do stuff!
