package main

import "fmt"

// Command interface (usado como "Observer")
type NotificationCommand interface {
	Execute(data string)
}

// Concrete Commands (como "Observers")
type EmailNotification struct{}
type SMSNotification struct{}
type PushNotification struct{}

func (e *EmailNotification) Execute(data string) {
	fmt.Println("Email notification:", data)
}

func (s *SMSNotification) Execute(data string) {
	fmt.Println("SMS notification:", data)
}

func (p *PushNotification) Execute(data string) {
	fmt.Println("Push notification:", data)
}

// Invoker (como "Subject")
type NotificationCenter struct {
	commands []NotificationCommand
}

func (nc *NotificationCenter) Register(command NotificationCommand) {
	nc.commands = append(nc.commands, command)
}

func (nc *NotificationCenter) NotifyAll(message string) {
	for _, cmd := range nc.commands {
		cmd.Execute(message)
	}
}

func main() {
	center := &NotificationCenter{}

	// Registrar "observers" como commands
	center.Register(&EmailNotification{})
	center.Register(&SMSNotification{})
	center.Register(&PushNotification{})

	// Notificar a todos (como Observer)
	center.NotifyAll("New message received!")
}
