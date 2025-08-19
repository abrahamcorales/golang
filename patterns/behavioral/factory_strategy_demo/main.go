package main

import "fmt"

// ===== FACTORY PATTERN =====
// Creates different types of payment processors

type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

type PayPalProcessor struct{}
type StripeProcessor struct{}
type CryptoProcessor struct{}

func (p PayPalProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("[PayPal] Processing $%.2f\n", amount)
	return nil
}

func (s StripeProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("[Stripe] Processing $%.2f\n", amount)
	return nil
}

func (c CryptoProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("[Crypto] Processing $%.2f\n", amount)
	return nil
}

// Factory function
func NewPaymentProcessor(provider string) (PaymentProcessor, error) {
	switch provider {
	case "paypal":
		return PayPalProcessor{}, nil
	case "stripe":
		return StripeProcessor{}, nil
	case "crypto":
		return CryptoProcessor{}, nil
	default:
		return nil, fmt.Errorf("unsupported provider: %s", provider)
	}
}

// ===== STRATEGY PATTERN =====
// Different pricing strategies for the same payment processor

type PricingStrategy interface {
	CalculatePrice(amount float64) float64
}

type StandardPricing struct{}
type PremiumPricing struct{}
type DiscountPricing struct{}

func (s StandardPricing) CalculatePrice(amount float64) float64 {
	return amount * 1.02 // 2% fee
}

func (p PremiumPricing) CalculatePrice(amount float64) float64 {
	return amount * 1.05 // 5% fee
}

func (d DiscountPricing) CalculatePrice(amount float64) float64 {
	return amount * 0.98 // 2% discount
}

// Context that uses both Factory and Strategy
type PaymentService struct {
	processor PaymentProcessor
	strategy  PricingStrategy
}

func NewPaymentService(provider string, pricingStrategy PricingStrategy) (*PaymentService, error) {
	processor, err := NewPaymentProcessor(provider)
	if err != nil {
		return nil, err
	}

	return &PaymentService{
		processor: processor,
		strategy:  pricingStrategy,
	}, nil
}

func (ps *PaymentService) ProcessPayment(amount float64) error {
	finalAmount := ps.strategy.CalculatePrice(amount)
	fmt.Printf("Original: $%.2f, Final: $%.2f\n", amount, finalAmount)
	return ps.processor.ProcessPayment(finalAmount)
}

func (ps *PaymentService) SetPricingStrategy(strategy PricingStrategy) {
	ps.strategy = strategy
}

func main() {
	fmt.Println("=== FACTORY + STRATEGY PATTERN EXAMPLE ===")

	// Factory: Create payment processor based on provider
	// Strategy: Use different pricing strategies

	// Example 1: PayPal with Standard pricing
	service1, _ := NewPaymentService("paypal", StandardPricing{})
	service1.ProcessPayment(100)

	// Example 2: Same PayPal processor, but with Premium pricing
	service1.SetPricingStrategy(PremiumPricing{})
	service1.ProcessPayment(100)

	// Example 3: Stripe with Discount pricing
	service2, _ := NewPaymentService("stripe", DiscountPricing{})
	service2.ProcessPayment(100)

	// Example 4: Switch pricing strategy at runtime
	service2.SetPricingStrategy(StandardPricing{})
	service2.ProcessPayment(100)
}
