# 01. Hello Go ☕

My very first Go program - because every programmer needs coffee!

## 📝 The Beginning - Simple Version

Let's start with the basics:
```go
package main

import "fmt"

func main() {
	fmt.Println("I want a cup of coffee!")
}
```

### Running It
```bash
go run main.go
```

**Output:**
```
I want a cup of coffee!
```

### What's Happening Here?

**`package main`**  
Every Go program starts here. This tells Go: "Hey, this is a program I can run!"

**`import "fmt"`**  
We're borrowing the `fmt` package to print stuff. It's like asking for tools before starting work.

**`func main()`**  
This is where your program starts. Think of it as the "Start" button.

**`fmt.Println()`**  
Prints text to the screen. Simple as that!

---

## 🔨 Building an Executable

Want to create a file you can run anytime?
```bash
go build main.go
```

This creates `main.exe` (or just `main` on Linux/Mac).

**Check what's in the folder:**
```bash
ls
```

**You'll see:**
```
Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
-a----        26.12.2025     00:42        2327552 main.exe
-a----        26.12.2025     00:07             91 main.go
```

**Now run it:**
```bash
./main.exe
```

**Output:**
```
I want a cup of coffee!
```

**Cool fact:** That `main.exe` is 2.3 MB! Why so big for 3 lines of code? Because Go includes EVERYTHING your program needs. No dependencies, no installations - just one file that works anywhere!

---

## 🎨 Making It Colorful

Plain text is boring. Let's add some color!

### Step 1: What Are External Packages?

Go comes with built-in packages (`fmt`, `os`, `time`, etc.) - these work out of the box.

But the community has created thousands of additional packages for extra features. Want colors? There's a package for that!

### Step 2: Finding the Right Package

I searched on [pkg.go.dev](https://pkg.go.dev/search?q=color) and found:

**`github.com/fatih/color`** - A popular package for colored console output!

- [Documentation](https://pkg.go.dev/github.com/fatih/color)
- [GitHub Repository](https://github.com/fatih/color)

### Step 3: Initialize Go Modules

Before installing external packages, we need to tell Go this is a proper project:
```bash
go mod init hello-go
```

**What happened?**  
A new file appeared: `go.mod`
```go
module hello-go

go 1.25.5
```

This file will track all the packages we use. Think of it as a shopping list for your project.

**Why is this needed?**  
Without modules, Go doesn't know where to save external packages. Modules organize everything!

### Step 4: Install the Color Package
```bash
go get github.com/fatih/color
```

**Look what happened:**
```
go: downloading github.com/fatih/color v1.18.0
go: downloading golang.org/x/sys v0.25.0
go: downloading github.com/mattn/go-colorable v0.1.13
go: downloading github.com/mattn/go-isatty v0.0.20
go: added github.com/fatih/color v1.18.0
go: added github.com/mattn/go-colorable v0.1.13
go: added github.com/mattn/go-isatty v0.0.20
go: added golang.org/x/sys v0.25.0
```

Go downloaded not just `color`, but also some helper packages it needs!

**Now check `go.mod`:**
```go
module hello-go

go 1.25.5

require (
	github.com/fatih/color v1.18.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.25.0 // indirect
)
```

All packages are listed! But what's `// indirect`?

**Direct vs Indirect:**
- **Direct:** Packages YOU use in your code
- **Indirect:** Packages that YOUR packages need (dependencies of dependencies)

Right now, all are marked `indirect` because we haven't used `color` in our code yet!

**Also, a new file appeared: `go.sum`**

This file contains checksums (security fingerprints) of all packages. It ensures nobody tampered with the code. Don't worry about editing it - Go manages it automatically!

### Step 5: Use the Color Package

Update `main.go`:
```go
package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("I want a cup of coffee!")
	color.Blue("I want a cup of coffee!")
}
```

**What changed?**
- Added `"github.com/fatih/color"` to imports
- Used `color.Blue()` to print blue text

### Step 6: Run It!
```bash
go run main.go
```

**Output:**
```
I want a cup of coffee!
I want a cup of coffee!  (← this one is BLUE! 🎨)
```

### Step 7: Clean Up Dependencies

Run this command:
```bash
go mod tidy
```

**What does it do?**
- Removes unused packages
- Adds missing packages
- Reorganizes `go.mod` for better readability

**Now check `go.mod` again:**
```go
module hello-go

go 1.25.5

require github.com/fatih/color v1.18.0

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.25.0 // indirect
)
```

**Notice the difference?**  
Now `color` is in the first `require` block (no `// indirect`)! Because we're DIRECTLY using it in our code.

The other packages are still indirect - they're needed by `color`, not by us.

**Always run `go mod tidy` before committing code!** It keeps everything clean.

---

## 📁 Final Project Structure
```
01. Hello Go/
├── main.go       # Your code
├── go.mod        # Tracks dependencies
├── go.sum        # Security checksums
└── main.exe      # Compiled program (after go build)
```

---

## 🧠 What I Learned

### Basic Go Structure
- ✅ Every program needs `package main`
- ✅ `import` brings in tools (packages)
- ✅ `func main()` is where execution starts
- ✅ `fmt.Println()` prints to console

### Building Programs
- ✅ `go run` - quick test (compile + run)
- ✅ `go build` - create executable file
- ✅ Go executables are self-contained (no external dependencies)

### Working with External Packages
- ✅ `go mod init` - start a new project
- ✅ `go get` - install external packages
- ✅ `go mod tidy` - clean up dependencies
- ✅ `go.mod` - tracks what packages you use
- ✅ `go.sum` - ensures package security

### Finding Packages
- ✅ [pkg.go.dev](https://pkg.go.dev) - search for packages
- ✅ Read package documentation before using
- ✅ Check GitHub stars and activity

### Direct vs Indirect Dependencies
- ✅ **Direct:** Packages you `import` in your code
- ✅ **Indirect:** Packages needed by your packages

---

## 🎯 Key Takeaway

**Go makes it easy to:**
1. Write simple programs with built-in packages
2. Extend functionality with external packages
3. Manage dependencies automatically
4. Build standalone executables

**One command to add features:** `go get <package>`  
**One command to clean up:** `go mod tidy`

Simple, powerful, organized! 🚀


---

*From plain text to colored output - day one success!* ☕🎨
