#

## Strategy Pattern Example (Go)

The Strategy pattern enables selecting an algorithm’s behavior at runtime. It defines a family of algorithms, encapsulates each one, and makes them interchangeable.

### Example: Payment Strategies

Suppose you have a shopping cart that can use different payment methods. With the Strategy pattern, you can switch between payment algorithms (like Credit Card or PayPal) at runtime without changing the cart’s code.

See the code in `strategy/main.go` for a runnable example.

## Observer Pattern Example (Go)

The Observer pattern defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically. This is useful for event-driven systems or implementing publish-subscribe mechanisms.

### Example: News Publisher and Subscribers

Suppose you have a news publisher that notifies multiple subscribers whenever a new article is published. With the Observer pattern, subscribers can register or unregister themselves, and the publisher notifies all registered subscribers of updates.

See the code in `observer/main.go` for a runnable example.