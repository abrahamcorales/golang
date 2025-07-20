## 🏃‍♂️ Go Interface & Composition Challenge

**Goal:**  
Practice Go interfaces and struct composition by modeling a simple animal sound system.

### Requirements

1. **Define an interface** called `Speaker` with a method `Speak() string`.
2. **Create two structs**: `Dog` and `Cat`, each implementing the `Speaker` interface.
   - `Dog.Speak()` should return `"Woof!"`
   - `Cat.Speak()` should return `"Meow!"`
3. **Create a struct** called `PetOwner` that has a field `Pet Speaker`.
4. **Write a function** `Announce(owner PetOwner)` that prints:  
   `"My pet says: <sound>"`  
   (where `<sound>` is the result of `owner.Pet.Speak()`)
5. **Write a main function** that:
   - Creates a `PetOwner` with a `Dog` and another with a `Cat`.
   - Calls `Announce` for both.

### Example Output

```
My pet says: Woof!
My pet says: Meow!
```

---

**Bonus:**  
- Add a third animal (e.g., `Parrot` that says `"Squawk!"`).
- Make `Announce` accept any type that implements `Speaker` (not just `PetOwner`).

## 🧩 Singleton Logger Challenge

**Goal:**  
Implement a thread-safe singleton logger in Go.

### Requirements

1. **Create a struct** called `Logger` with a method `Log(message string)` that prints the message prefixed with `[LOG]:`.
2. **Implement a function** `GetLogger()` that always returns the same instance of `Logger` (singleton).
   - Ensure your implementation is thread-safe (use `sync.Once`).
3. **Write a main function** that:
   - Gets the singleton logger instance twice.
   - Calls `Log` with two different messages using both instances.
   - Demonstrates that both instances are the same (e.g., by comparing their addresses or using `==`).

### Example Output

```
[LOG]: First message
[LOG]: Second message
Logger instances are the same: true
```

---

**Bonus:**  
- Add a field to `Logger` that counts the number of log messages, and print the count after logging.