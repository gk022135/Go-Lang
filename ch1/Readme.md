# GO Lang — One Shot Notes (Quick Revision)

> _Beginner-friendly · Hinglish mein explain · Interview-ready_

---

## Why Go Lang?

> **"Most procedural languages allow complexity. Go actively prevents it."**

Go lang **highly concurrent** aur **performance-efficient** hai. Iska design philosophy simple hai — jo cheez zaroori nahi, woh dali hi nahi.

### Go mein YEH nahi hai (intentionally!)
| Feature | Why removed? |
|---|---|
| No Inheritance | Composition prefer karo, inheritance nahi |
| No Macros | Code readable rehna chahiye |
| No Pointer Arithmetic | Memory bugs avoid karne ke liye |
| No Overloading | Confusion avoid karne ke liye |

### Go kyun use karein?
1. **Simple syntax** — C jaisi speed, Python jaisi readability
2. **Built-in concurrency** — Goroutines + Channels (BIG reason!)
3. **Fast compilation** — Seconds mein compile
4. **Single binary output** — Deploy karna super easy
5. **Standard library** — Bahut powerful built-in tools

```go
// Traditional thread vs Go goroutine
// C threads → complex, heavy
// Java threads → heavier, boilerplate
// Go goroutine → bas ek 'go' keyword!
go func() {
    fmt.Println("Hello from goroutine")
}()
```

---

## Variables in Go

### Types
| Type | Description |
|---|---|
| `var` | Mutable — value baad mein change kar sakte ho |
| `const` | Immutable — ek baar set, phir change nahi |

```go
// var — change ho sakta hai
var name = "Gaurav"
name = "Gaurav2"   // allowed

// const — FIXED rehta hai
const PI = 3.14
PI = 3.0        // ERROR! const change nahi hota

// Short declaration (function ke andar hi)
age := 25       // var age = 25 ka shortcut
```

### Scope
```go
// Package level (global)
var globalVar = "I am global"
const AppName = "MyApp"

func main() {
    // Function level (local)
    var localVar = "I am local"
    shortVar := "also local"
}
```

---

## Symbols in Go — Meaning

| Symbol | Name | Use |
|---|---|---|
| `:=` | Short Assignment | Variable declare + assign karo (local only) |
| `==` | Equality Check | Do values equal hain? |
| `!=` | Not Equal | Equal nahi hain? |
| `=` | Assignment | Pehle se declared var mein value daalo |
| `&` | AND / Address-of | Bitwise AND ya pointer address |
| `\|` | OR | Bitwise OR |
| `^` | XOR | Bitwise XOR |
| `&&` | Logical AND | Dono true hone chahiye |
| `\|\|` | Logical OR | Koi ek true ho |
| `!` | NOT | Opposite |
| `<-` | Channel Operator | Goroutine channels mein data bhejna/lena |
| `...` | Variadic / Spread | Multiple args ya slice spread |

```go
x := 10          // declare and assign
x = 20           // update (already declared)
y == x           // compare (returns bool)
ch <- value      // channel mein value daalo
value := <-ch    // channel se value nikalo
```

---

## Data Types in Go

### Primitives
```go
// Integer
var age int = 25
var bigNum int64 = 9999999999

// Float
var price float32 = 99.99
var precise float64 = 3.14159265

// String
var name string = "Gopher"

// Bool
var isActive bool = true
```

### Arrays — Fixed Size
```go
// Simple array
var marks [5]int = [5]int{90, 85, 92, 78, 88}

// Short syntax
b := [5]int{1, 2, 3, 4, 5}

// 2D Array
var matrix [3][3]int
c := [3][3]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}
```

### Slices — Dynamic Size (Arrays se zyada use hota hai!)
```go
// Slice — growable array
fruits := []string{"apple", "banana", "mango"}

// append karo
fruits = append(fruits, "orange")

// make se
nums := make([]int, 5)   // length 5
```

### Maps — Key-Value Store
```go
// Declaration
m := make(map[string]int)

// Initialize
m["a"] = 1
m["b"] = 2
m["c"] = 3

// One-liner
n := map[string]int{"x": 10, "y": 20, "z": 30}

// Value check karo
val, exists := m["a"]
if exists {
    fmt.Println("Value:", val)
}

// Delete
delete(m, "a")
```

### Ranges — Loop with Index+Value
```go
arr := [5]int{1, 2, 3, 4, 5}

// Array range
for i, v := range arr {
    fmt.Println("arr[", i, "] =", v)
}

// Slice range
fruits := []string{"apple", "banana"}
for i, fruit := range fruits {
    fmt.Println(i, fruit)
}

// Map range
scores := map[string]int{"Alice": 90, "Bob": 85}
for name, score := range scores {
    fmt.Println(name, "->", score)
}
```

### Functions
```go
// Basic function
func add(a int, b int) int {
    return a + b
}

// Multiple return values (Go ka superpower!)
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

// Variadic function
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

### Structs (Go ka "Class")
```go
type Person struct {
    Name string
    Age  int
}

// Use
p := Person{Name: "Rahul", Age: 25}
fmt.Println(p.Name)  // Rahul

// Method on struct
func (p Person) Greet() string {
    return "Hello, I am " + p.Name
}
```

---

## Control Flow

```go
// If-else
if x > 10 {
    fmt.Println("big")
} else if x == 10 {
    fmt.Println("equal")
} else {
    fmt.Println("small")
}

// For loop (Go mein sirf for loop hai, while nahi!)
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// While jaisa
for condition {
    // ...
}

// Infinite loop
for {
    // break se bahar niklo
}

// Switch
switch day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("Weekend coming!")
default:
    fmt.Println("Normal day")
}
```

---

## ⚡ CONCURRENCY IN GO — Full Deep Dive

> **Yeh section most important hai! Interviews mein sabse zyada pucha jaata hai.**

---

### Concurrency kya hai? (Simple explanation)

Socho tum ek **chef** ho restaurant kitchen mein:

**Sequential (Ek kaam ek baar):**
```
Pasta banao → wait karo → Pizza banao → wait karo → Salad banao
Total time: 30 min + 20 min + 10 min = 60 minutes
```

**Concurrent (Saath mein kaam):**
```
Pasta flame pe daalo → Pizza oven mein daalo → Salad kaato
Saare saath chal rahe hain!
Total time: max(30, 20, 10) = 30 minutes
```

**Concurrency = Multiple tasks ko manage karna, taaki CPU idle na baithe.**

> Concurrency ≠ Parallelism!
> - **Concurrency**: Multiple tasks *manage* karna (single CPU bhi kar sakta hai)
> - **Parallelism**: Multiple tasks *literally ek saath* run karna (multiple CPUs chahiye)

---

## GOROUTINES — Go ka Superpower

### Goroutine kya hai?

**Normal thread** — ek heavy worker. 1MB+ memory, OS manage karta hai.

**Goroutine** — ek lightweight worker. Sirf **2KB** se start! Go runtime manage karta hai.

```
OS Threads:    [████████ 1MB] [████████ 1MB] [████████ 1MB]
Goroutines:    [▪2KB] [▪2KB] [▪2KB] [▪2KB] [▪2KB] [▪2KB]

Tum ek machine pe thousands of goroutines chala sakte ho!
```

### Goroutine kaise shuru karta hai?

Bas `go` keyword lagao function ke pehle — ho gaya!

```go
package main

import (
    "fmt"
    "time"
)

func bolो(message string) {
    fmt.Println(message)
}

func main() {
    // Normal call — main ruk jaata hai
    bolo("Main thread se namaste")

    // Goroutine — background mein chala do, main wait nahi karega
    go bolo("Goroutine se namaste")

    // Wait karo warna main exit ho jaayega goroutine se pehle
    time.Sleep(1 * time.Second)
}
```

### Problem without sync:

```go
func main() {
    go fmt.Println("Goroutine")
    // main yahan se exit ho gaya!
    // goroutine run hi nahi hua
}
// Output: (kuch nahi!)
```

### Solution — WaitGroup (Proper tarika)

```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()  // kaam khatam hone pe Done() call karo
    fmt.Printf("Worker %d kaam kar raha hai\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)           // ek aur goroutine add kiya
        go worker(i, &wg)   // goroutine shuru karo
    }

    wg.Wait()  // sab goroutines khatam hone ka wait karo
    fmt.Println("Sab kaam ho gaya!")
}

// Output (order vary kar sakta hai):
// Worker 3 kaam kar raha hai
// Worker 1 kaam kar raha hai
// Worker 5 kaam kar raha hai
// Worker 2 kaam kar raha hai
// Worker 4 kaam kar raha hai
// Sab kaam ho gaya!
```

---

## CHANNELS — Goroutines ke beech baat karna

> **"Do not communicate by sharing memory; instead, share memory by communicating."**
> — Go team ka philosophy

Channel ek **pipe** hai jisse goroutines data safely share karte hain.

```go
// Channel banana
ch := make(chan int)          // unbuffered
ch := make(chan int, 5)       // buffered (5 tak store kar sakta hai)

// Data bhejna
ch <- 42        // channel mein 42 daalo

// Data lena
val := <-ch     // channel se value nikalo
```

### Simple Channel Example:

```go
package main

import "fmt"

func sum(nums []int, ch chan int) {
    total := 0
    for _, n := range nums {
        total += n
    }
    ch <- total  // result channel mein bhejo
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    ch := make(chan int)

    // Pehla half
    go sum(nums[:5], ch)   // 1+2+3+4+5 = 15
    // Doosra half
    go sum(nums[5:], ch)   // 6+7+8+9+10 = 40

    // Dono results lo
    x, y := <-ch, <-ch

    fmt.Println("Total:", x+y)  // 55
}
```

---

## 🔧 PIPELINE PATTERN — Go ka Real-World Concurrency

> **Pipeline = Assembly line ki tarah!**

Socho car factory:
```
[Frame banao] → [Engine daalo] → [Paint karo] → [QA check]
     Stage 1   →    Stage 2    →    Stage 3   →   Stage 4
```

Har stage ek **goroutine** hai. Data ek se doosre mein **channels** ke through jaata hai.

---

### Pipeline Example — Numbers ko Process karna

**Problem**: Numbers 1-5 lo, unhe double karo, phir square karo.

```go
package main

import "fmt"

// Stage 1: Generator — numbers produce karo
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n  // channel mein bhejo
        }
        close(out)  // kaam khatam, channel band karo
    }()
    return out
}

// Stage 2: Double karo
func double(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {  // jab tak channel open hai
            out <- n * 2
        }
        close(out)
    }()
    return out
}

// Stage 3: Square karo
func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func main() {
    // Pipeline connect karo!
    //  1,2,3,4,5  →  double  →  square  →  print
    nums := generate(1, 2, 3, 4, 5)
    doubled := double(nums)
    squared := square(doubled)

    // Results print karo
    for result := range squared {
        fmt.Println(result)
    }
}

// Output:
// 4    (1 → double=2 → square=4)
// 16   (2 → double=4 → square=16)
// 36   (3 → double=6 → square=36)
// 64   (4 → double=8 → square=64)
// 100  (5 → double=10 → square=100)
```

**Yeh kya hua step by step:**
```
Input:   1  →  2  →  3  →  4  →  5
Double:  2  →  4  →  6  →  8  →  10
Square:  4  → 16  → 36  → 64  → 100
```

---

### Fan-Out, Fan-In Pattern (Advanced Pipeline)

**Problem**: Ek source se data lo, multiple workers mein distribute karo (speed ke liye), phir results combine karo.

```
              ┌─→ Worker 1 ─┐
Source ───────┼─→ Worker 2 ─┼──→ Result Collector
              └─→ Worker 3 ─┘
   Fan-Out ↑                  ↑ Fan-In
```

```go
package main

import (
    "fmt"
    "sync"
)

// Source: jobs produce karo
func producer(jobs []int) <-chan int {
    out := make(chan int)
    go func() {
        for _, j := range jobs {
            out <- j
        }
        close(out)
    }()
    return out
}

// Worker: job lo, process karo (yahaan square kar rahe hain)
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for j := range jobs {
        result := j * j
        fmt.Printf("Worker %d: %d ka square = %d\n", id, j, result)
        results <- result
    }
}

// Collector: sab results collect karo
func collector(results <-chan int) {
    total := 0
    for r := range results {
        total += r
    }
    fmt.Println("\nTotal sum of squares:", total)
}

func main() {
    jobs := producer([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
    results := make(chan int, 10)
    var wg sync.WaitGroup

    // 3 workers shuru karo (Fan-Out)
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, jobs, results, &wg)
    }

    // Jab sab workers done ho jaayein to results close karo
    go func() {
        wg.Wait()
        close(results)
    }()

    // Results collect karo (Fan-In)
    collector(results)
}
```

---

### Select — Multiple Channels pe sunna

`select` = goroutines ke liye `switch` statement

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "Channel 1 se message"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "Channel 2 se message"
    }()

    // Jo pehle ready ho, woh lo
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Mila:", msg1)
        case msg2 := <-ch2:
            fmt.Println("Mila:", msg2)
        case <-time.After(3 * time.Second):
            fmt.Println("Timeout! Koi response nahi")
        }
    }
}
```

---

## Goroutine vs Thread — Quick Comparison

| Feature | OS Thread | Goroutine |
|---|---|---|
| Memory | ~1MB | ~2KB |
| Startup time | Slow | Very fast |
| Managed by | OS | Go Runtime |
| Switch cost | Expensive | Very cheap |
| Max on machine | ~1,000s | ~100,000s |
| Communication | Shared memory (risky) | Channels (safe) |

---

## Race Condition aur Mutex

**Race condition** = Do goroutines ek saath same data change karein — dangerous!

```go
// BAD — Race condition!
var counter int
for i := 0; i < 1000; i++ {
    go func() {
        counter++  // multiple goroutines ek saath counter++ karein
    }()
}
// counter ka value unpredictable hoga!

// GOOD — Mutex use karo
var mu sync.Mutex
var counter int
for i := 0; i < 1000; i++ {
    go func() {
        mu.Lock()    // lock lo
        counter++
        mu.Unlock()  // lock chodo
    }()
}
```

---

## Quick Reference Cheatsheet

```
GOROUTINE    → go funcName()
CHANNEL      → ch := make(chan Type)
SEND         → ch <- value
RECEIVE      → val := <-ch
CLOSE        → close(ch)
WAITGROUP    → var wg sync.WaitGroup → wg.Add(1) → wg.Done() → wg.Wait()
SELECT       → multiple channels monitor karna
MUTEX        → mu.Lock() / mu.Unlock() — race condition se bachao
PIPELINE     → goroutines ko channels se chain karo
```

---

## Real World Use Cases

| Pattern | Kab use karein |
|---|---|
| Goroutine | Background tasks, API calls, file processing |
| Pipeline | Data transformation, ETL, image processing |
| Fan-Out/Fan-In | Parallel processing, web scraping, batch jobs |
| Select + Timeout | Network calls with timeout |
| Mutex | Shared counter/map updates |

---

## ⚡ Summary — Go Lang in 5 Points

1. **Simple** — Minimal syntax, maximum clarity
2. **Fast** — Compiled language, near C performance
3. **Concurrent** — Goroutines + Channels = easy concurrency
4. **Safe** — Garbage collected, no pointer arithmetic
5. **Scalable** — Thousands of goroutines, minimal memory

---

> **Yaad rakho**: Go mein concurrency ka mantra hai —
> **"Don't communicate by sharing memory, share memory by communicating."**
> Channels use karo, mutex sirf tab jab bahut zaroori ho.

---

_Happy Coding! 🐹_