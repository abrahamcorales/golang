package main

import (
	"fmt"
)

type Ordered interface {
	~int | ~float64 | ~string
}

func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(Min(3, 7))            // 3
	fmt.Println(Min(2.5, 1.2))        // 1.2
	fmt.Println(Min("go", "generic")) // generic
}
