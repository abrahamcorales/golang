# 🧱 Module 2 – Creational Design Patterns in Go

This module covers the three most common creational patterns in Go:

- Factory (with and without interfaces)
- Singleton (thread-safe with `sync.Once`)
- Builder (for complex structs with optional fields)

---

## 🔧 Factory Pattern

### 🤔 Why use interfaces?

Using interfaces in the Factory pattern decouples the object creation logic from the code that uses those objects. This makes your code easier to extend, test, and maintain, since you can change the concrete implementation without modifying the rest of your code. Without interfaces, your code becomes more rigid and tightly coupled to specific implementations, reducing flexibility and making it harder to scale or adapt.

### ✅ Purposer

Encapsulates object creation logic to avoid tight coupling and enable flexible implementations.

### 🔸 Without Interface

```go
type ClientType string

const (
	HTTP ClientType = "http"
	GRPC ClientType = "grpc"
)

type Client struct {
	Protocol string
}

func NewClient(t ClientType) *Client {
	switch t {
	case HTTP:
		return &Client{Protocol: "HTTP/1.1"}
	case GRPC:
		return &Client{Protocol: "gRPC"}
	default:
		return &Client{Protocol: "Unknown"}
	}
}
```

**Uso:**
```go
client := NewClient(HTTP)
fmt.Println(client.Protocol) // Output: HTTP/1.1

client2 := NewClient(GRPC)
fmt.Println(client2.Protocol) // Output: gRPC

client3 := NewClient("other")
fmt.Println(client3.Protocol) // Output: Unknown
```

### 🔸 WIth Interface

```go

type Storage interface {
	Save(data string) error
}

type S3Storage struct{}
func (s S3Storage) Save(data string) error {
	fmt.Println("Saving to S3:", data)
	return nil
}

type DynamoDBStorage struct{}
func (d DynamoDBStorage) Save(data string) error {
	fmt.Println("Saving to DynamoDB:", data)
	return nil
}

func NewStorage(kind string) Storage {
	switch kind {
	case "s3":
		return S3Storage{}
	case "dynamo":
		return DynamoDBStorage{}
	default:
		return nil
	}
}
```

**Uso:**
```go
storage := NewStorage("s3")
if storage != nil {
	storage.Save("hello world")
}
```

## 🧩 Singleton Pattern

### ✅ Purpose
Ensures a class has only one instance and provides a global point of access to it.

### 🔸 Thread-safe Singleton with sync.Once

```go
import (
	"fmt"
	"sync"
)

type singleton struct {
	Data string
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{Data: "I am the only instance"}
	})
	return instance
}
```

**Usage:**
```go
s1 := GetInstance()
s2 := GetInstance()
fmt.Println(s1 == s2) // Output: true
fmt.Println(s1.Data)  // Output: I am the only instance
```

---

## 🏗️ Builder Pattern

### ✅ Purpose
Helps construct complex objects step by step, especially when many optional fields are involved.

### 🔸 Example

```go
type User struct {
	Name    string
	Age     int
	Address string
	Email   string
}

type UserBuilder struct {
	user User
}

func NewUserBuilder(name string) *UserBuilder {
	return &UserBuilder{user: User{Name: name}}
}

func (b *UserBuilder) Age(age int) *UserBuilder {
	b.user.Age = age
	return b
}

func (b *UserBuilder) Address(address string) *UserBuilder {
	b.user.Address = address
	return b
}

func (b *UserBuilder) Email(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) Build() User {
	return b.user
}
```

**Usage:**
```go
user := NewUserBuilder("Alice").Age(30).Email("alice@example.com").Build()
fmt.Printf("%+v\n", user)
// Output: {Name:Alice Age:30 Address: Email:alice@example.com}
```

---


