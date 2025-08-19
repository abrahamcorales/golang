# Go Generics

## What are Generics?

Generics in Go are a way to write functions, types, and data structures that can work with multiple data types without having to rewrite the same code for each specific type. They allow you to write more flexible and reusable code by using type parameters.

## Why Use Generics?

Before generics were introduced in Go 1.18, you had to either:
- Write separate functions for each data type
- Use `interface{}` (which loses type safety)
- Use code generation tools

Generics solve these problems by providing:
- **Type Safety**: Compile-time type checking
- **Code Reuse**: Write once, use with many types
- **Performance**: No runtime type assertions needed
- **Readability**: Clear intent in the code

## Basic Syntax

### Type Parameters

Generics use square brackets `[]` to declare type parameters:

```go
func Min[T comparable](a, b T) T {
    if a < b {
        return a
    }
    return b
}
```

Here, `T` is a type parameter that must satisfy the `comparable` constraint.

### Constraints

Constraints limit what types can be used:

```go
// Numeric constraint
func Sum[T ~int | ~float64](numbers []T) T {
    var sum T
    for _, n := range numbers {
        sum += n
    }
    return sum
}

// Interface constraint
func Stringify[T fmt.Stringer](item T) string {
    return item.String()
}
```

## Common Use Cases

### 1. Generic Functions

```go
// Works with any comparable type
func Contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

// Usage
numbers := []int{1, 2, 3, 4, 5}
names := []string{"Alice", "Bob", "Charlie"}

fmt.Println(Contains(numbers, 3))  // true
fmt.Println(Contains(names, "David")) // false
```

### 2. Generic Types

```go
// Generic stack implementation
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
    if len(s.items) == 0 {
        var zero T
        return zero, errors.New("stack is empty")
    }
    
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, nil
}

// Usage
intStack := &Stack[int]{}
intStack.Push(1)
intStack.Push(2)

stringStack := &Stack[string]{}
stringStack.Push("hello")
stringStack.Push("world")
```

### 3. Generic Data Structures

```go
// Generic map with default values
type DefaultMap[K comparable, V any] struct {
    data map[K]V
    defaultValue V
}

func NewDefaultMap[K comparable, V any](defaultValue V) *DefaultMap[K, V] {
    return &DefaultMap[K, V]{
        data: make(map[K]V),
        defaultValue: defaultValue,
    }
}

func (dm *DefaultMap[K, V]) Get(key K) V {
    if value, exists := dm.data[key]; exists {
        return value
    }
    return dm.defaultValue
}

func (dm *DefaultMap[K, V]) Set(key K, value V) {
    dm.data[key] = value
}
```

## Built-in Constraints

Go provides several useful constraints:

- **`any`**: Any type (equivalent to `interface{}`)
- **`comparable`**: Types that can be compared with `==` and `!=`
- **`Ordered`**: Types that can be ordered with `<`, `<=`, `>`, `>=`
- **`Integer`**: Integer types
- **`Float`**: Floating-point types
- **`Complex`**: Complex number types

## Custom Constraints

You can define your own constraints:

```go
// Custom constraint for numeric types
type Number interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
    ~float32 | ~float64
}

func Average[T Number](numbers []T) T {
    if len(numbers) == 0 {
        var zero T
        return zero
    }
    
    var sum T
    for _, n := range numbers {
        sum += n
    }
    return sum / T(len(numbers))
}
```

## Best Practices

1. **Use meaningful type parameter names**: `T`, `K`, `V` for simple cases, descriptive names for complex ones
2. **Keep constraints minimal**: Only require what you actually need
3. **Consider performance**: Generics are compiled to specific types, so there's no runtime overhead
4. **Document your constraints**: Make it clear what types your generic code expects

## Limitations

- Generics cannot be used with methods (only functions and types)
- Type parameters cannot be used in method receivers
- Some complex type relationships are not expressible

## When to Use Generics

Use generics when:
- You need to work with multiple types in the same way
- You want type safety without code duplication
- You're building reusable libraries or utilities

Don't use generics when:
- The code is only used with one specific type
- The abstraction makes the code harder to understand
- You need runtime type flexibility

## Example: Complete Generic Utility

```go
package main

import (
    "fmt"
    "sort"
)

// Generic utility functions
func Reverse[T any](slice []T) []T {
    result := make([]T, len(slice))
    for i, j := 0, len(slice)-1; i < len(slice); i, j = i+1, j-1 {
        result[i] = slice[j]
    }
    return result
}

func Filter[T any](slice []T, predicate func(T) bool) []T {
    var result []T
    for _, item := range slice {
        if predicate(item) {
            result = append(result, item)
        }
    }
    return result
}

func Map[T, U any](slice []T, transform func(T) U) []U {
    result := make([]U, len(slice))
    for i, item := range slice {
        result[i] = transform(item)
    }
    return result
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Filter even numbers
    evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
    fmt.Println("Even numbers:", evens)
    
    // Double each number
    doubled := Map(numbers, func(n int) int { return n * 2 })
    fmt.Println("Doubled:", doubled)
    
    // Reverse the slice
    reversed := Reverse(numbers)
    fmt.Println("Reversed:", reversed)
}
```

Generics make Go more expressive and powerful while maintaining the language's simplicity and performance characteristics. They're particularly useful for writing libraries and utilities that need to work with multiple data types.