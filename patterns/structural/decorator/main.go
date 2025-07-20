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
type Milk struct {
	CoffeeDecorator
}

func (m *Milk) Cost() int {
	return m.Coffee.Cost() + 2
}
func (m *Milk) Ingredients() string {
	return m.Coffee.Ingredients() + ", Milk"
}

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
