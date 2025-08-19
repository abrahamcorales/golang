package main

import "fmt"

// Strategy
type PaymentStrategy interface {
	Pay(amount float64)
}

// Concrete Strategies
type CreditCard struct {
	Name, CardNumber string
}

func (c *CreditCard) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using Credit Card (%s)\n", amount, c.CardNumber)
}

type PayPal struct {
	Email string
}

func (p *PayPal) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using PayPal (%s)\n", amount, p.Email)
}

// Context
type ShoppingCart struct {
	Payment PaymentStrategy
}

func (s *ShoppingCart) Checkout(amount float64) {
	s.Payment.Pay(amount)
}

func main() {
	cart := &ShoppingCart{}

	cart.Payment = &CreditCard{Name: "Alice", CardNumber: "1234-5678"}
	cart.Checkout(50.0) // Paid $50.00 using Credit Card (1234-5678)

	cart.Payment = &PayPal{Email: "alice@example.com"}
	cart.Checkout(25.0) // Paid $25.00 using PayPal (alice@example.com)

	// Exercise implementation
	fmt.Println("\n=== SHIPPING STRATEGY EXERCISE ===")

	shipping := &ShippingContext{}

	// Test different shipping strategies
	shipping.Strategy = &StandardShipping{}
	shipping.Ship("Laptop", 100)

	shipping.Strategy = &ExpressShipping{}
	shipping.Ship("Laptop", 100)

	shipping.Strategy = &OvernightShipping{}
	shipping.Ship("Laptop", 100)

	// Exercise 2: Simple Greeting Strategy
	fmt.Println("\n=== SIMPLE GREETING STRATEGY ===")

	greeter := &GreetingContext{}

	// Test different greeting strategies
	greeter.Strategy = &FormalGreeting{}
	greeter.Greet()

	greeter.Strategy = &CasualGreeting{}
	greeter.Greet()

	greeter.Strategy = &FriendlyGreeting{}
	greeter.Greet()
}

/*
EXERCISE: Shipping Strategy

Create a shipping strategy system where you can choose different shipping methods.

TODO: Implement the following:

1. Create a ShippingStrategy interface with method:
  - Ship(item string, distance int)

2. Create concrete strategies:
  - StandardShipping: "Standard delivery in 5-7 days"
  - ExpressShipping: "Express delivery in 2-3 days"
  - OvernightShipping: "Overnight delivery"

3. Create a ShippingContext struct that uses ShippingStrategy

4. Test with different shipping methods for the same package

Expected Output:
Package: "Laptop", Distance: 100km
Standard: Standard delivery in 5-7 days
Express: Express delivery in 2-3 days
Overnight: Overnight delivery
*/

type ShippingStrategy interface {
	Ship(item string, distance int)
}

type ShippingContext struct {
	Strategy ShippingStrategy
}

func (s *ShippingContext) Ship(item string, distance int) {
	s.Strategy.Ship(item, distance)
}

type StandardShipping struct{}

func (s *StandardShipping) Ship(item string, distance int) {
	fmt.Printf("Standard: Standard delivery in 5-7 days\n")
}

type ExpressShipping struct{}

func (e *ExpressShipping) Ship(item string, distance int) {
	fmt.Printf("Express: Express delivery in 2-3 days\n")
}

type OvernightShipping struct{}

func (o *OvernightShipping) Ship(item string, distance int) {
	fmt.Printf("Overnight: Overnight delivery\n")
}

/*
EXERCISE 2: Simple Greeting Strategy

Create a simple greeting strategy system.

TODO: Implement the following:

1. Create a GreetingStrategy interface with method:
   - Greet(name string)

2. Create concrete strategies:
   - FormalGreeting: "Hello, Mr. {name}"
   - CasualGreeting: "Hi {name}!"
   - FriendlyGreeting: "Hey {name}, how are you?"

3. Create a GreetingContext struct that uses GreetingStrategy

4. Test with different greeting styles for the same person

Expected Output:
Formal: Hello, Mr. John
Casual: Hi John!
Friendly: Hey John, how are you?
*/

type GreetingStrategy interface {
	Greet()
}

type GreetingContext struct {
	Strategy GreetingStrategy
}

func (g *GreetingContext) Greet() {}

type FormalGreeting struct {
}

func (g *FormalGreeting) Greet() {}

type CasualGreeting struct {
}

func (g *CasualGreeting) Greet() {}

type FriendlyGreeting struct {
}

func (g *FriendlyGreeting) Greet() {}
