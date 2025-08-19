## Adapter Pattern Example (Go)

The Adapter pattern allows incompatible interfaces to work together. It acts as a bridge between two incompatible interfaces by wrapping an existing class with a new interface.

### Example: Media Player with Different Audio Formats

```go
package main

import "fmt"

// Target Interface - What the client expects
type MediaPlayer interface {
	Play(audioType string, fileName string)
}

// Adaptee - Existing class with incompatible interface
type AdvancedMediaPlayer interface {
	PlayVlc(fileName string)
	PlayMp4(fileName string)
}

// Concrete Adaptee implementations
type VlcPlayer struct{}

func (v *VlcPlayer) PlayVlc(fileName string) {
	fmt.Printf("Playing vlc file: %s\n", fileName)
}

func (v *VlcPlayer) PlayMp4(fileName string) {
	// VlcPlayer doesn't support MP4
}

type Mp4Player struct{}

func (m *Mp4Player) PlayVlc(fileName string) {
	// Mp4Player doesn't support VLC
}

func (m *Mp4Player) PlayMp4(fileName string) {
	fmt.Printf("Playing mp4 file: %s\n", fileName)
}

// Adapter - Makes incompatible interfaces work together
type MediaAdapter struct {
	advancedPlayer AdvancedMediaPlayer
}

func (m *MediaAdapter) Play(audioType string, fileName string) {
	if audioType == "vlc" {
		m.advancedPlayer.PlayVlc(fileName)
	} else if audioType == "mp4" {
		m.advancedPlayer.PlayMp4(fileName)
	}
}

// Concrete Adapter implementations
type VlcAdapter struct {
	MediaAdapter
}

func NewVlcAdapter() *VlcAdapter {
	return &VlcAdapter{
		MediaAdapter: MediaAdapter{
			advancedPlayer: &VlcPlayer{},
		},
	}
}

type Mp4Adapter struct {
	MediaAdapter
}

func NewMp4Adapter() *Mp4Adapter {
	return &Mp4Adapter{
		MediaAdapter: MediaAdapter{
			advancedPlayer: &Mp4Player{},
		},
	}
}

// Client - Uses the target interface
type AudioPlayer struct{}

func (a *AudioPlayer) Play(audioType string, fileName string) {
	if audioType == "mp3" {
		fmt.Printf("Playing mp3 file: %s\n", fileName)
	} else if audioType == "vlc" || audioType == "mp4" {
		var adapter MediaPlayer
		if audioType == "vlc" {
			adapter = NewVlcAdapter()
		} else {
			adapter = NewMp4Adapter()
		}
		adapter.Play(audioType, fileName)
	} else {
		fmt.Printf("Invalid media type: %s\n", audioType)
	}
}

func main() {
	player := &AudioPlayer{}
	
	player.Play("mp3", "song.mp3")
	player.Play("vlc", "movie.vlc")
	player.Play("mp4", "video.mp4")
	player.Play("avi", "movie.avi")
}
```

### Exercise: Payment Gateway Adapter

Create an adapter to make different payment gateways work with a unified interface.

```go
// TODO: Implement the following:

// 1. Create a PaymentGateway interface (target):
//    - ProcessPayment(amount float64) bool

// 2. Create existing payment systems (adaptees):
//    - PayPal: PayWithPayPal(email string, amount float64) bool
//    - Stripe: ChargeWithStripe(cardNumber string, amount float64) bool

// 3. Create adapters:
//    - PayPalAdapter: wraps PayPal to match PaymentGateway interface
//    - StripeAdapter: wraps Stripe to match PaymentGateway interface

// 4. Test with different payment methods

// Expected Output:
// PayPal payment processed: $50.00
// Stripe payment processed: $100.00
```

This example shows how you can make incompatible payment systems work together through adapters.