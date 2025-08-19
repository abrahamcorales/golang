package main

// 1. Create a PaymentCard interface with methods:
//    - GetAnnualFee() int
//    - GetFeatures() string

type PaymentCard interface {
	GetAnnualFee() int
	GetFeatures() string
}

// 2. Create a BasicCard struct that implements PaymentCard
//    - Annual Fee: 0
//    - Features: "Basic Payment"

type BasiCard struct {
}

func (b *BasiCard) GetAnnualFee() int {
	return 0
}

func (b *BasiCard) GetFeatures() string {
	return "Features: Basic Payment"
}

type CardDecorator struct {
	PaymentCard
}
type Rewards struct {
	CardDecorator
}

func (c *Rewards) GetAnnualFee() int {
	return 50
}
func (c *Rewards) GetFeatures() string {
	return "Features: Basic Payment"
}

func (c *CardDecorator) GetAnnualFee() int {
	return 50
}
func (c *CardDecorator) GetFeatures() string {
	return "Features: Basic Payment"
}

func main() {

}

// 3. Create a CardDecorator struct that embeds PaymentCard

// 4. Create decorators for card features:
//    - Rewards (fee: +50, features: ", Cashback Rewards")
//    - Travel (fee: +100, features: ", Travel Insurance")
//    - Premium (fee: +200, features: ", Premium Support")

// 5. Test your implementation:
//    - Start with BasicCard
//    - Add Rewards
//    - Add Travel
//    - Add Premium
//    - Print annual fee and features at each step
