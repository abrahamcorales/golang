package main

import "fmt"

// Command Interface
type Command interface {
	Execute()
	Undo()
}

// Concrete Commands
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.TurnOn()
}

func (c *LightOnCommand) Undo() {
	c.light.TurnOff()
}

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.TurnOff()
}

func (c *LightOffCommand) Undo() {
	c.light.TurnOn()
}

// Receiver
type Light struct {
	isOn bool
}

func (l *Light) TurnOn() {
	l.isOn = true
	fmt.Println("Light is ON")
}

func (l *Light) TurnOff() {
	l.isOn = false
	fmt.Println("Light is OFF")
}

func (l *Light) GetStatus() string {
	if l.isOn {
		return "ON"
	}
	return "OFF"
}

// Invoker
type RemoteControl struct {
	commands []Command
	history  []Command
}

func (rc *RemoteControl) SetCommand(command Command) {
	rc.commands = append(rc.commands, command)
}

func (rc *RemoteControl) PressButton(index int) {
	if index < len(rc.commands) {
		rc.commands[index].Execute()
		rc.history = append(rc.history, rc.commands[index])
	}
}

func (rc *RemoteControl) UndoLast() {
	if len(rc.history) > 0 {
		lastCommand := rc.history[len(rc.history)-1]
		lastCommand.Undo()
		rc.history = rc.history[:len(rc.history)-1]
	}
}

func main() {
	// Create receiver
	light := &Light{}

	// Create commands
	lightOn := &LightOnCommand{light: light}
	lightOff := &LightOffCommand{light: light}

	// Create invoker
	remote := &RemoteControl{}

	// Set commands
	remote.SetCommand(lightOn)  // Button 0
	remote.SetCommand(lightOff) // Button 1

	fmt.Println("=== COMMAND PATTERN DEMO ===")
	fmt.Printf("Light status: %s\n", light.GetStatus())

	// Execute commands
	fmt.Println("\nPressing button 0 (Turn ON):")
	remote.PressButton(0)
	fmt.Printf("Light status: %s\n", light.GetStatus())

	fmt.Println("\nPressing button 1 (Turn OFF):")
	remote.PressButton(1)
	fmt.Printf("Light status: %s\n", light.GetStatus())

	// Undo last command
	fmt.Println("\nUndoing last command:")
	remote.UndoLast()
	fmt.Printf("Light status: %s\n", light.GetStatus())

	// Undo again
	fmt.Println("\nUndoing last command:")
	remote.UndoLast()
	fmt.Printf("Light status: %s\n", light.GetStatus())
}
