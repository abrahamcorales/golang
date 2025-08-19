# Command Pattern vs Observer Pattern

## Command Pattern

The **Command pattern** encapsulates a request as an object, allowing you to parameterize clients with different requests, queue operations, and support undo operations.

### Key Components:
- **Command Interface**: Defines `Execute()` and `Undo()` methods
- **Concrete Commands**: Implement specific actions
- **Receiver**: The object that performs the actual work
- **Invoker**: Triggers the command execution

### Use Cases:
- **Remote controls** (like TV remotes)
- **Undo/Redo functionality** in text editors
- **Macro recording** and playback
- **Queue operations** for background processing

### Example:
```go
// Command encapsulates a request
type Command interface {
    Execute()
    Undo()
}

// Invoker doesn't know what command it executes
remote.PressButton(0) // Could be any command
```

---

## Observer Pattern

The **Observer pattern** defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified automatically.

### Key Components:
- **Subject**: Maintains list of observers and notifies them
- **Observer Interface**: Defines update method
- **Concrete Observers**: React to subject's state changes

### Use Cases:
- **Event-driven systems** (GUI frameworks)
- **Publish-subscribe** mechanisms
- **Model-View** architecture
- **Notification systems**

### Example:
```go
// Subject notifies all observers when state changes
estacion.Notificar("Tormenta eléctrica") // All alerts receive this
```

---

## Key Differences

| Aspect | Command Pattern | Observer Pattern |
|--------|----------------|------------------|
| **Purpose** | Encapsulate requests as objects | Notify multiple objects of state changes |
| **Relationship** | One-to-one (Invoker → Command) | One-to-many (Subject → Observers) |
| **Direction** | Invoker → Receiver | Subject → Observers |
| **Timing** | On-demand execution | Automatic notification |
| **State** | Commands can be queued/undone | Observers react to state changes |
| **Coupling** | Invoker decoupled from Receiver | Subject loosely coupled to Observers |

## When to Use Which?

### Use Command Pattern when:
- You need **undo/redo functionality**
- You want to **queue operations**
- You need to **parameterize objects** with operations
- You want to **log operations** for audit trails

### Use Observer Pattern when:
- You need **event-driven architecture**
- Multiple objects need to **react to state changes**
- You want **loose coupling** between subject and observers
- You're building **notification systems**

## Summary

- **Command**: "Do this specific action" (encapsulated request)
- **Observer**: "Something changed, notify everyone" (event notification)

Both patterns promote loose coupling but serve different purposes in software design. 