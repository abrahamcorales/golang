## 🏋️‍♂️ Factory Pattern Challenge

**Goal:**  
Implement a simple factory in Go that creates different types of notifications (e.g., Email, SMS) based on a string input.

### Requirements

1. **Define an interface** called `Notifier` with a method `Send(message string) string`.
2. **Create two structs**: `EmailNotifier` and `SMSNotifier`, each implementing the `Notifier` interface.
   - `EmailNotifier.Send` should return: `"Email sent: <message>"`
   - `SMSNotifier.Send` should return: `"SMS sent: <message>"`
3. **Write a factory function** `NewNotifier(kind string) Notifier` that:
   - Returns an `EmailNotifier` if `kind` is `"email"`
   - Returns an `SMSNotifier` if `kind` is `"sms"`
   - Returns `nil` for any other value
4. **Write a main function** that:
   - Uses the factory to create both types of notifiers and calls their `Send` method with a sample message.
   - Prints the results.

### Example Output

```
Email sent: Hello, Factory!
SMS sent: Hello, Factory!
```

---

**Bonus:**  
- Add a third notifier type (e.g., `PushNotifier`).
- Handle the case where the factory returns `nil` gracefully in `main`.