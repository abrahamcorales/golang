# 🧱 SOLID Principles in Go

This README explains how to apply the **SOLID principles** in the Go programming language.  
Go does not support inheritance, so it uses **interfaces**, **composition**, and **clean code design** to achieve the same goals.

---

## ✅ 1. Single Responsibility Principle (SRP)

> A struct or function should have **only one reason to change**.

In Go:
- Keep structs small and focused.
- Each interface should represent a single responsibility.

### 🧪 Example

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 1. Business logic only
type Invoice struct {
	ID     string
	Amount float64
	Tax    float64
}

func (i *Invoice) Total() float64 {
	return i.Amount * (1 + i.Tax)
}

// 2. Separate component for saving
type InvoiceSaver struct{}

func (s InvoiceSaver) SaveToFile(i *Invoice, filename string) error {
	data, err := json.Marshal(i)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func main() {
	inv := Invoice{ID: "INV-001", Amount: 100, Tax: 0.2}
	fmt.Println("Total:", inv.Total())

	saver := InvoiceSaver{}
	if err := saver.SaveToFile(&inv, "invoice.json"); err != nil {
		fmt.Println("Error saving invoice:", err)
	}
}
```
## ✅ 2. Open/Closed Principle (OCP)

### 📖 Definition
**Software entities should be open for extension, but closed for modification.**  
This means that you should be able to add new functionality without changing existing working code.

---

### 💡 How to apply it in Go

Go doesn't support inheritance, so we follow OCP by:
- ✅ Using **interfaces** to define behaviors
- ✅ Passing different implementations as dependencies (composition)
- ❌ Avoiding method overriding (not supported in Go)

```go

type Notifier interface {
	Notify(message string) error
}

type EmailNotifier struct{}
func (e EmailNotifier) Notify(msg string) error {
	fmt.Println("Sending email:", msg)
	return nil
}

type SlackNotifier struct{}
func (s SlackNotifier) Notify(msg string) error {
	fmt.Println("Sending Slack message:", msg)
	return nil
}
func SendAlert(n Notifier) {
	n.Notify("Server down")
}

func main() {
	var notifier Notifier

	// Switch between notifiers without changing SendAlert
	notifier = EmailNotifier{}
	SendAlert(notifier)

	notifier = SlackNotifier{}
	SendAlert(notifier)
}
```
## ✅ 3. Liskov Substitution Principle (LSP)

### 📖 Definition
**If type `S` is a subtype of type `T`, we should be able to substitute `T` with `S` without breaking the program.**  
In Go, if a struct implements an interface, it must behave consistently with the expected contract.

---

### 💡 How to apply it in Go

- ✅ Keep interfaces small and focused
- ✅ Ensure implementations do not surprise the caller
- ❌ Avoid side effects or unexpected behavior in concrete types

---

### ✅ Example

```go
type Notifier interface {
	Notify(message string) error
}

type EmailNotifier struct{}
func (e EmailNotifier) Notify(msg string) error {
	fmt.Println("Email:", msg)
	return nil
}

type SlackNotifier struct{}
func (s SlackNotifier) Notify(msg string) error {
	fmt.Println("Slack:", msg)
	return nil
}

// This function should work with *any* Notifier
func SendSystemAlert(n Notifier) {
	n.Notify("System is down!")
}

func main() {
	var n Notifier

	n = EmailNotifier{}
	SendSystemAlert(n)

	n = SlackNotifier{}
	SendSystemAlert(n)
}
```

## ✅ 4. Interface Segregation Principle (ISP)

### 📖 Definition
**Clients should not be forced to depend on interfaces they do not use.**  
Instead of having large interfaces, split them into small, specific ones.

---

### 💡 How to apply it in Go

Go naturally encourages small interfaces. You should:
- ✅ Create minimal interfaces like `io.Reader`, `io.Writer`
- ✅ Design by **capability** (e.g., if your function only needs `.Read()`, accept a `Reader`)
- ❌ Avoid "God interfaces" that force implementations to satisfy unnecessary methods

---

### ✅ Example

```go
// Bad - Fat Interface
type File interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Close() error
}

// Good - Segregated Interfaces
type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// Client that only needs to read
func ProcessData(r Reader) {
	buf := make([]byte, 100)
	r.Read(buf)
	fmt.Println("Read data")
}
```

## ✅ 5. Dependency Inversion Principle (DIP)

### 📖 Definition
**High-level code (like business logic) should depend on interfaces, not on specific implementations.**  
We pass those implementations using constructors or struct fields (**dependency injection**).

---

### 💡 How to apply it in Go

- ✅ Define interfaces at the consumer level (not in the low-level package)
- ✅ Inject dependencies via constructors or struct fields
- ✅ Decouple business logic from infrastructure concerns

---

### ✅ Example

```go
// Define the abstraction that high-level code depends on
type Storage interface {
	Save(data string) error
}

// Low-level implementation
type FileStorage struct{}
func (fs FileStorage) Save(data string) error {
	fmt.Println("Saving to file:", data)
	return nil
}

// High-level logic depends on the abstraction, not the concrete type
type ReportService struct {
	storage Storage
}

// Dependency is injected via constructor
func NewReportService(s Storage) *ReportService {
	return &ReportService{storage: s}
}

func (r *ReportService) GenerateReport() {
	fmt.Println("Generating report...")
	r.storage.Save("report content")
}

func main() {
	fs := FileStorage{}
	service := NewReportService(fs)
	service.GenerateReport()
}
```
