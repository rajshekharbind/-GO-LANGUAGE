# 🚀 GO-LANGUAGE – Complete Golang Learning Repository

This repository is a **structured, beginner-to-advanced Golang learning collection**.
Each folder focuses on **one core Go concept**, with **clean examples**, **practical use-cases**, and **hands-on code**.

The goal of this repository is to:

* Build **strong Go fundamentals**
* Learn **idiomatic Go practices**
* Prepare for **interviews, backend development, and system programming**

---

## 📂 Project Structure

```
GO-LANGUAGE/
│
├── arr/
├── crud/
├── data-conversion/
├── defer/
├── error/
├── file/
├── for-loop/
├── function/
├── goroutine/
├── if-else/
├── inputstring/
├── json/
├── map/
├── myutil/
├── pointer/
├── slice/
├── struct/
├── switch/
├── sync_waitgroup/
├── time/
├── URL/
├── webRequest/
│
├── example.txt
├── go.mod
└── main.go
```

---

## 🧠 Topics Covered (Detailed Explanation)

### 📌 1. arr (Arrays)

* Fixed-size collections in Go
* Declaration and initialization
* Iteration techniques
* Memory behavior

**Key Concepts:**

* Index-based access
* Value semantics
* Compile-time size

---

### 📌 2. slice (Slices)

* Dynamic arrays built on top of arrays
* `append`, `len`, `cap`
* Slice reallocation

**Why Important?**
Slices are the **most commonly used collection** in Go.

---

### 📌 3. map (Maps)

* Key-value data structures
* Insert, update, delete
* Iteration over maps

**Key Notes:**

* Reference type
* Fast lookups
* Unordered

---

### 📌 4. function

* Function declaration & invocation
* Multiple return values
* Named return values
* Variadic functions

---

### 📌 5. pointer

* Pointer basics
* `&` and `*` operators
* Pass-by-value vs pass-by-reference

**Why Pointers Matter?**

* Performance
* Shared state

---

### 📌 6. struct

* Custom data types
* Struct fields
* Nested structs
* Struct methods

---

### 📌 7. if-else

* Conditional logic
* Short variable declarations
* Best practices

---

### 📌 8. switch

* Switch cases
* Multiple case values
* Type switch

---

### 📌 9. for-loop

* Classic for-loop
* While-style loop
* Infinite loop
* Range-based loop

---

### 📌 10. inputstring

* Reading input from user
* `fmt.Scan`, `Scanln`
* Buffer-based input

---

### 📌 11. data-conversion

* Type casting
* String ↔ int conversion
* Float conversions

---

### 📌 12. json

* Encoding structs to JSON
* Decoding JSON into structs
* Working with APIs

---

### 📌 13. error

* Error handling in Go
* Custom errors
* `errors.New` and `fmt.Errorf`

**Go Philosophy:**

> Errors are values

---

### 📌 14. file

* File creation
* Read & write operations
* File permissions

---

### 📌 15. defer

* Deferred function calls
* Stack behavior
* Resource cleanup

---

### 📌 16. goroutine

* Lightweight concurrency
* `go` keyword
* Non-blocking execution

---

### 📌 17. sync_waitgroup

* Synchronizing goroutines
* `sync.WaitGroup`
* `Add`, `Done`, `Wait`

**Used for:**

* Parallel processing
* Safe goroutine completion

---

### 📌 18. time

* Time delays
* Timers
* Measuring execution time

---

### 📌 19. URL

* URL parsing
* Query parameters
* URL encoding & decoding

---

### 📌 20. webRequest

* HTTP GET requests
* Handling responses
* Reading API data

---

### 📌 21. crud

* Create, Read, Update, Delete
* REST-style operations
* JSON-based handling

---

## 📄 Other Files

### 📘 example.txt

* Sample data file
* Used for file-handling demos

### 📘 main.go

* Entry point of the application
* Used for testing or combined execution

### 📘 go.mod

* Module definition file
* Dependency management

---

## 🛠 Requirements

* Go 1.20+
* Basic programming knowledge

---

## ▶️ How to Run

```bash
go run main.go
```

Or run individual topic files:

```bash
go run goroutine/main.go
```

---

## 🎯 Learning Outcomes

* Strong Golang fundamentals
* Concurrency mastery
* Clean and idiomatic Go code
* Interview-ready concepts

---

## ⭐ Who Should Use This Repo?

* Beginners learning Go
* Backend developers
* Students & interview candidates
* Anyone revising Go concepts

---

## 🙌 Contribution

Feel free to fork, improve examples, or add advanced topics like:

* Channels
* Mutex
* Context
* Testing

---

## 📌 Author

**Rajshekhar**
Happy Coding with Go 🐹🚀
