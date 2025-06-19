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
```
