package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
}

func main() {
	count := GetCounter()
	count.Increment()
	count.Increment()
	fmt.Println("Counter value:", count.value)
}

var (
	Once     sync.Once
	instance *Counter
)

func GetCounter() *Counter {
	Once.Do(func() {
		instance = &Counter{}
	})
	return instance
}
func (c *Counter) Increment() int {
	c.value++
	return c.value

}
