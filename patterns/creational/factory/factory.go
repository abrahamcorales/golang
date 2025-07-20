package factory

import (
	"fmt"
	"os"
)

func main() {
	provider := os.Getenv("PAYMENT_PROVIDER")

	procesor, _ := NewPaymentProcessor(provider)

	_ = procesor.ProcessPayment(63)
}

type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}
type PayPalProcessor struct{}

func (p PayPalProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("[PayPal] Payment of $%.2f processed successfully.\n", amount)
	return nil
}

type StripeProcessor struct{}

func (s StripeProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("[Stripe] Payment of $%.2f processed successfully.\n", amount)
	return nil
}

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
