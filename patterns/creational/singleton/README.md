## 🧩 Singleton Pattern Challenge

**Goal:**  
Implement a thread-safe singleton in Go that manages a simple configuration value.

### Requirements

1. **Create a struct** called `Config` with a field `AppName string`.
2. **Implement a function** `GetConfig()` that always returns the same instance of `Config` (singleton).
   - Ensure your implementation is thread-safe (use `sync.Once`).
3. **Add a method** to `Config` called `SetAppName(name string)` to update the `AppName` field.
4. **Write a main function** that:
   - Gets the singleton instance twice.
   - Sets the `AppName` using one instance.
   - Prints the `AppName` from both instances to show they are the same object.

### Example Output

```
AppName from config1: MyApp
AppName from config2: MyApp
```

---

**Bonus:**  
- Prevent direct instantiation of `Config` from outside the package (hint: use unexported struct or constructor).