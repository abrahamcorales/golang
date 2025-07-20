# Builder Pattern in Go

The **Builder Pattern** is a creational design pattern that lets you construct complex objects step by step. It's useful when an object needs to be created with many optional fields or when the construction process is complex.

## When to Use

- When you have a struct with many fields, some of which are optional.
- When you want to avoid a constructor with many parameters.
- When you want to make object creation readable and maintainable.

## Example

See `builder.go` for a complete example.

```go
user := NewUserBuilder().
    Name("Alice").
    Age(30).
    Email("alice@example.com").
    Address("123 Main St").
    Build()
```

## Exercise

**Task:**  
Implement a `CarBuilder` that can build a `Car` struct with the following fields:

- `Brand` (string)
- `Model` (string)
- `Year` (int)
- `Color` (string)
- `Electric` (bool)

**Requirements:**

1. Create a `Car` struct and a `CarBuilder` struct.
2. Implement builder methods for each field (e.g., `Brand()`, `Model()`, etc.).
3. Add a `Build()` method that returns the constructed `Car`.
4. In `main()`, use your builder to create at least two different cars and print them.

**Bonus:**  
Make the builder chainable (return `*CarBuilder` from each method).

---

**Happy coding!** 