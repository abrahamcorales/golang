package main

import "fmt"

// Component - Interface base
type Text interface {
	Display() string
}

// Concrete Component - Implementación básica
type SimpleText struct {
	Content string
}

func (s *SimpleText) Display() string {
	return s.Content
}

// Decorator Base - Envuelve el componente
type TextDecorator struct {
	Text
}

// Concrete Decorators - Agregan funcionalidad
type BoldDecorator struct {
	TextDecorator
}

func (b *BoldDecorator) Display() string {
	return "**" + b.Text.Display() + "**"
}

type ItalicDecorator struct {
	TextDecorator
}

func (i *ItalicDecorator) Display() string {
	return "*" + i.Text.Display() + "*"
}

type UnderlineDecorator struct {
	TextDecorator
}

func (u *UnderlineDecorator) Display() string {
	return "__" + u.Text.Display() + "__"
}

func main() {
	// Texto básico
	var text Text = &SimpleText{Content: "Hello World"}
	fmt.Println("Original:", text.Display())

	// Agregar negrita
	text = &BoldDecorator{TextDecorator{text}}
	fmt.Println("Bold:", text.Display())

	// Agregar cursiva
	text = &ItalicDecorator{TextDecorator{text}}
	fmt.Println("Bold + Italic:", text.Display())

	// Agregar subrayado
	text = &UnderlineDecorator{TextDecorator{text}}
	fmt.Println("Bold + Italic + Underline:", text.Display())

	var sandwich Sandwich = &BasicSandwich{}
	fmt.Println(sandwich.GetDescription()) // Bread

	sandwich = &Lettuce{SandwichDecorator{sandwich}}
	fmt.Println(sandwich.GetDescription()) // Bread, Lettuce

	sandwich = &Tomato{SandwichDecorator{sandwich}}
	fmt.Println(sandwich.GetDescription()) // Bread, Lettuce, Tomato

}

/*
EJERCICIO: Sandwich Decorator

Crea un sistema de decoradores para agregar ingredientes a un sandwich.

TODO: Implementa lo siguiente:

1. Crea una interfaz Sandwich con método:
  - GetDescription() string

2. Crea un componente concreto BasicSandwich:
  - Descripción: "Bread"

3. Crea un decorador base SandwichDecorator

4. Crea decoradores concretos:
  - Lettuce (descripción: ", Lettuce")
  - Tomato (descripción: ", Tomato")
  - Cheese (descripción: ", Cheese")

5. Prueba agregando ingredientes uno por uno

Salida esperada:
Bread
Bread, Lettuce
Bread, Lettuce, Tomato
Bread, Lettuce, Tomato, Cheese
*/
type Sandwich interface {
	GetDescription() string
}

type BasicSandwich struct{}

func (b *BasicSandwich) GetDescription() string {
	return "Bread"
}

type SandwichDecorator struct {
	Sandwich
}

type Lettuce struct {
	SandwichDecorator
}

func (l *Lettuce) GetDescription() string {
	return l.Sandwich.GetDescription() + ", Lettuce"
}

type Tomato struct {
	SandwichDecorator
}

func (t *Tomato) GetDescription() string {
	return t.Sandwich.GetDescription() + ", Tomato"
}
