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
}
