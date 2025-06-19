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
