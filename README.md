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
type Logger struct{}

func (l Logger) Log(msg string) {
	fmt.Println(msg)
}

type UserService struct {
	logger Logger
}

func (u UserService) CreateUser(name string) {
	// Logic for creating a user
	u.logger.Log("User created: " + name)
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
```
