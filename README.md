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

## 🔍 Key Differences Between SOLID Principles in Go

| Principle | Focus                    | Question it Answers                                              | Risk if Violated                                                                 |
|-----------|--------------------------|------------------------------------------------------------------|----------------------------------------------------------------------------------|
| **SRP** (Single Responsibility) | Separation of concerns       | "Does this component do more than one thing?"                    | Code becomes harder to read, test, or change without side effects               |
| **OCP** (Open/Closed)          | Extending functionality      | "How can I add a new behavior without changing existing code?"   | You end up rewriting stable code to add features                                |
| **LSP** (Liskov Substitution) | Replacing implementation     | "Will my code still work if I swap one implementation for another?" | Replacements break existing logic or introduce bugs                             |
| **ISP** (Interface Segregation) | Keeping interfaces focused    | "Am I forcing implementations to satisfy unused methods?"         | You couple components unnecessarily and reduce flexibility                      |
| **DIP** (Dependency Inversion) | Decoupling modules           | "Is my high-level logic depending on low-level details?"          | Code becomes tightly coupled, hard to test, and inflexible to change            |


# SOLID Principles vs Design Patterns

| Name                          | Type        | Category              | Goal / Purpose                                                                 | Example Use Case                                        |
|-------------------------------|-------------|------------------------|--------------------------------------------------------------------------------|----------------------------------------------------------|
| **S**: Single Responsibility  | Principle   | SOLID (Class Design)  | A class should have only one reason to change                                  | Separate logging from business logic                     |
| **O**: Open/Closed            | Principle   | SOLID (Extensibility) | Software entities should be open for extension, but closed for modification    | Use interfaces to allow new behaviors                    |
| **L**: Liskov Substitution    | Principle   | SOLID (Inheritance)   | Derived classes must be substitutable for their base classes                   | Replace a parent class with any child without breaking   |
| **I**: Interface Segregation  | Principle   | SOLID (Abstraction)   | Clients should not be forced to depend on interfaces they don't use            | Split large interfaces into smaller, role-based ones     |
| **D**: Dependency Inversion   | Principle   | SOLID (Decoupling)    | Depend on abstractions, not concrete implementations                           | Inject services using interfaces                         |
| **Factory**                   | Pattern     | Creational            | Encapsulate object creation logic                                              | Create `DatabaseClient` based on config                  |
| **Singleton**                 | Pattern     | Creational            | Ensure a class has only one instance                                           | Centralized configuration loader                         |
| **Builder**                   | Pattern     | Creational            | Construct complex objects step-by-step                                         | Fluent API to build a large struct                      |
| **Strategy**                  | Pattern     | Behavioral            | Define a family of algorithms, encapsulate them, and make them interchangeable | Payment strategies: `PayPal`, `CreditCard`, etc.         |
| **Observer**                  | Pattern     | Behavioral            | Notify multiple objects when state changes                                     | Event-driven systems, pub-sub                            |
| **Decorator**                 | Pattern     | Structural            | Add responsibilities to objects dynamically                                    | Add logging or caching to an existing service            |
| **Adapter**                   | Pattern     | Structural            | Make incompatible interfaces work together                                     | Wrap an old API to fit a new interface                   |
| **Singleflight**              | Pattern     | Concurrency           | Prevent duplicate function calls for the same key                              | Deduplicate concurrent cache misses or HTTP requests     |
| **Non-blocking Cache**        | Pattern     | Concurrency           | Return stale or default value while fresh data is recomputed asynchronously    | Improve responsiveness under high load                   |
| **Fan-out / Parallelism**     | Pattern     | Concurrency           | Launch multiple workers in parallel to speed up batch processing               | Parallel API calls or I/O-heavy operations               |
| **Worker Pool**               | Pattern     | Concurrency           | Limit concurrency by distributing work across a pool of fixed-size workers     | Limit DB access concurrency, bulk email sending          |
| **Pipeline**                  | Pattern     | Concurrency           | Chain processing stages using channels                                         | ETL jobs, streaming data processing                      |
| **Semaphore**                 | Pattern     | Concurrency           | Limit the number of concurrent goroutines                                      | Restrict simultaneous DB connections                     |
| **Context Cancellation**      | Pattern     | Concurrency           | Propagate cancellation signals to goroutines                                   | Cancel HTTP requests, shutdown background jobs           |
| **Ticker/Timer**              | Pattern     | Concurrency           | Schedule periodic or delayed tasks                                             | Heartbeats, scheduled cache refresh                      |
| **Select Statement**          | Pattern     | Concurrency           | Multiplex channel operations, handle multiple events                           | Wait for multiple channel inputs or timeouts             |
| **Future/Promise**            | Pattern     | Concurrency           | Represent a value that will be available in the future                         | Async computation results                                |
| **Rate Limiter**              | Pattern     | Concurrency           | Control the rate of events or requests                                         | API request throttling                                   |
