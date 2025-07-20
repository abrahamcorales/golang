package main

import "fmt"

// Observer interface
type Subscriber interface {
	Update(article string)
}

// Concrete Observer
type EmailSubscriber struct {
	Email string
}

func (e *EmailSubscriber) Update(article string) {
	fmt.Printf("Email to %s: New article published: %s\n", e.Email, article)
}

// Concrete Observer
type SmsSubscriber struct {
	Phone string
}

func (s *SmsSubscriber) Update(article string) {
	fmt.Printf("SMS to %s: New article published: %s\n", s.Phone, article)
}

// Subject (Publisher)
type Publisher struct {
	subscribers []Subscriber
}

func (p *Publisher) Register(sub Subscriber) {
	p.subscribers = append(p.subscribers, sub)
}
func (p *Publisher) Unregister(sub Subscriber) {
	for i, s := range p.subscribers {
		if s == sub {
			p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
			break
		}
	}
}
func (p *Publisher) Notify(article string) {
	for _, sub := range p.subscribers {
		sub.Update(article)
	}
}

func main() {
	publisher := &Publisher{}

	emailSub := &EmailSubscriber{Email: "alice@example.com"}
	smsSub := &SmsSubscriber{Phone: "+1234567890"}

	publisher.Register(emailSub)
	publisher.Register(smsSub)

	publisher.Notify("Observer Pattern in Go")
	// Output:
	// Email to alice@example.com: New article published: Observer Pattern in Go
	// SMS to +1234567890: New article published: Observer Pattern in Go

	publisher.Unregister(emailSub)
	publisher.Notify("Another Article")
	// Output:
	// SMS to +1234567890: New article published: Another Article
}
