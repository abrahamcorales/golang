package main

import (
	"fmt"
)

type Car struct {
	Brand    string
	Model    string
	Year     int
	Color    string
	Electric bool
}

func NewCarBuilder() *Car {
	return &Car{}
}

func (c *Car) WithBrand(name string) *Car {
	c.Brand = name
	return c
}

func (c *Car) WithModel(model string) *Car {
	c.Model = model
	return c
}

func (c *Car) WithYear(year int) *Car {
	c.Year = year
	return c
}

func (c *Car) WithColor(color string) *Car {
	c.Color = color
	return c
}

func (c *Car) WithElectric(electric bool) *Car {
	c.Electric = electric
	return c
}

func (c *Car) Build() Car {
	return *c
}

func main() {

	car := NewCarBuilder().
		WithBrand("Ford").
		WithModel("Mustang").
		WithYear(2024).
		WithColor("Red").
		WithElectric(false).
		Build()
	fmt.Println(car)

}
