package main

import (
	"fmt"
)

func main() {
    a, b := 100.1, 1000.34

	fmt.Printf("Before:\na = %v\nb = %v\n\n", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("After:\na = %v\nb = %v\n", a, b)
}