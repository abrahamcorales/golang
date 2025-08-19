# Factory Pattern vs Strategy Pattern

## Factory Pattern

The **Factory pattern** is a **creational** design pattern that encapsulates object creation logic. It provides a way to create objects without specifying their exact class.

### Key Components:
- **Factory Method**: Creates objects based on input parameters
- **Product Interface**: Common interface for all created objects
- **Concrete Products**: Specific implementations

### Purpose:
- **Object Creation**: Centralize and standardize object instantiation
- **Configuration-based**: Create objects based on configuration or parameters
- **Abstraction**: Hide complex object creation logic

### Example from your code:
```go
// Factory creates objects based on configuration
func NewPaymentProcessor(provider string) (PaymentProcessor, error) {
    switch provider {
    case "paypal":
        return PayPalProcessor{}, nil
    case "stripe":
        return StripeProcessor{}, nil
    default:
        return nil, fmt.Errorf("unsupported payment provider: %s", provider)
    }
}

// Usage
processor, _ := NewPaymentProcessor("paypal")
processor.ProcessPayment(100)
```

---

## Strategy Pattern

The **Strategy pattern** is a **behavioral** design pattern that defines a family of algorithms, encapsulates each one, and makes them interchangeable at runtime.

### Key Components:
- **Strategy Interface**: Defines the algorithm contract
- **Concrete Strategies**: Different algorithm implementations
- **Context**: Uses the strategy and can change it at runtime

### Purpose:
- **Algorithm Selection**: Choose different algorithms at runtime
- **Behavior Switching**: Change behavior without changing the context
- **Extensibility**: Add new algorithms without modifying existing code

### Example from your code:
```go
// Context can switch strategies at runtime
type ShoppingCart struct {
    Payment PaymentStrategy
}

func (s *ShoppingCart) Checkout(amount float64) {
    s.Payment.Pay(amount)
}

// Usage - switching strategies
cart := &ShoppingCart{}
cart.Payment = &CreditCard{Name: "Alice", CardNumber: "1234-5678"}
cart.Checkout(50.0)

cart.Payment = &PayPal{Email: "alice@example.com"}
cart.Checkout(25.0)
```

---

## Key Differences

| Aspect | Factory Pattern | Strategy Pattern |
|--------|----------------|------------------|
| **Category** | Creational | Behavioral |
| **Purpose** | Object creation | Algorithm selection |
| **Timing** | Compile-time/Configuration | Runtime |
| **Flexibility** | Fixed creation logic | Dynamic behavior switching |
| **Usage** | Create objects once | Switch behaviors frequently |
| **Dependency** | Configuration/Parameters | Runtime decisions |

## When to Use Which?

### Use Factory Pattern when:
- You need to **create objects** based on configuration
- Object creation logic is **complex**
- You want to **centralize** object creation
- Objects are created **once** and used throughout the application

### Use Strategy Pattern when:
- You need to **switch algorithms** at runtime
- Multiple **behaviors** can be applied to the same context
- You want to **extend** functionality without modifying existing code
- **Behavior changes** frequently during execution

## Real-world Analogy

### Factory Pattern:
Think of a **car factory**:
- You order a car by specifying the model ("sedan", "SUV")
- The factory creates the appropriate car based on your specification
- Once created, the car doesn't change its type

### Strategy Pattern:
Think of a **GPS navigation**:
- You can switch between different routing strategies ("fastest", "shortest", "scenic")
- The GPS context remains the same, but the routing algorithm changes
- You can switch strategies multiple times during a trip

## Summary

- **Factory**: "Create this type of object for me" (object creation)
- **Strategy**: "Use this algorithm for this behavior" (behavior selection)

Both patterns promote loose coupling but serve different purposes in software design.