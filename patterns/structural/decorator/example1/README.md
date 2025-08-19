## Decorator Pattern Example (Go)

The Decorator pattern allows you to add new functionality to an object dynamically, without altering its structure. This is achieved by wrapping the original object with a new object.

### Example: Coffee with Decorators

```go
package main

import "fmt"

// Component
type Coffee interface {
    Cost() int
    Ingredients() string
}

// Concrete Component
type SimpleCoffee struct{}

func (s *SimpleCoffee) Cost() int {
    return 5
}
func (s *SimpleCoffee) Ingredients() string {
    return "Coffee"
}

// Decorator Base
type CoffeeDecorator struct {
    Coffee
}

// Concrete Decorators
type Milk struct{ CoffeeDecorator }
func (m *Milk) Cost() int           { return m.Coffee.Cost() + 2 }
func (m *Milk) Ingredients() string { return m.Coffee.Ingredients() + ", Milk" }

type Sugar struct{ CoffeeDecorator }
func (s *Sugar) Cost() int           { return s.Coffee.Cost() + 1 }
func (s *Sugar) Ingredients() string { return s.Coffee.Ingredients() + ", Sugar" }

func main() {
    var c Coffee = &SimpleCoffee{}
    fmt.Println(c.Cost(), c.Ingredients()) // 5 Coffee

    c = &Milk{CoffeeDecorator{c}}
    fmt.Println(c.Cost(), c.Ingredients()) // 7 Coffee, Milk

    c = &Sugar{CoffeeDecorator{c}}
    fmt.Println(c.Cost(), c.Ingredients()) // 8 Coffee, Milk, Sugar
}
```

This example shows how you can add Milk and Sugar to a Coffee dynamically, each time wrapping the previous object.

### Exercise: Payment Card Features

Create a Payment Card decorator system where you can add features dynamically.

```go
// TODO: Implement the following structure:

// 1. Create a PaymentCard interface with methods:
//    - GetAnnualFee() int
//    - GetFeatures() string

// 2. Create a BasicCard struct that implements PaymentCard
//    - Annual Fee: 0
//    - Features: "Basic Payment"

// 3. Create a CardDecorator struct that embeds PaymentCard

// 4. Create decorators for card features:
//    - Rewards (fee: +50, features: ", Cashback Rewards")
//    - Travel (fee: +100, features: ", Travel Insurance")
//    - Premium (fee: +200, features: ", Premium Support")

// 5. Test your implementation:
//    - Start with BasicCard
//    - Add Rewards
//    - Add Travel
//    - Add Premium
//    - Print annual fee and features at each step
```

**Expected Output:**
```
Annual Fee: $0, Features: Basic Payment
Annual Fee: $50, Features: Basic Payment, Cashback Rewards
Annual Fee: $150, Features: Basic Payment, Cashback Rewards, Travel Insurance
Annual Fee: $350, Features: Basic Payment, Cashback Rewards, Travel Insurance, Premium Support
```