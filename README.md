# 🎫 Ticket Booking CLI (Go Project)

A simple command-line ticket booking system built using **Go (Golang)**.  
This project demonstrates basic Go concepts like variables, slices, loops, user input, validation, and string handling.

---

## 🚀 Features

- Displays available tickets
- Takes user input (name & email)
- Basic user validation (name length, email format)
- Tracks bookings using slices
- Adds new bookings dynamically
- Greets returning users
- Uses loops and conditionals
- Demonstrates Go fundamentals

---

## 🧠 Concepts Used

- Variables & Constants
- `fmt` package (input/output)
- Slices (`[]string`)
- `append()` function
- `for range` loop
- `if-else` conditions
- String functions (`strings.Contains`)
- Pointers (`&variable`)
- Type checking (`%T`)
- Basic input validation

---

## 📦 Project Structure

```
main.go
```

---

## ▶️ How to Run

### 1. Install Go
```bash
go version
```

---

### 2. Run the program
```bash
go run main.go
```

OR build and run:

```bash
go build main.go
./main
```

(on Windows)
```bash
main.exe
```

---

## 🧪 Sample Output

```
Remaining tickets: 100 tickets available.
Hello, World!
Welcome John to Go programming.
Please Enter Your Name:-Alex
Your name is: Alex
Please Enter Your Email Id:-alex@gmail.com
Your email id is: alex@gmail.com
The first booking is: Nana
[Nana Praxy Jimmy]
3
[Nana Praxy Jimmy Alex]
Welcome back, Alex!
```

---

## ⚠️ Validations

### Name:
- Must be at least 2 characters

### Email:
- Must contain `@`

---

## 📚 What I Learned

- Go basics and syntax
- Slices and dynamic arrays
- User input handling
- Conditionals and loops
- String validation
- CLI application structure

---

## 🔥 Future Improvements

- Seat selection system
- File/database storage
- Multiple events support
- REST API version (backend project)
- Goroutines for concurrency

---

## 👨‍💻 Author Prabuddha Pal

Built while learning Go 🚀
