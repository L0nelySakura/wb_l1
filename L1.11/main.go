package main

import (
	"fmt"
)

func main() {
	A := []int{1, 2, 3, 3}
	B := []int{2, 3, 3, 4}

	set := make(map[int]bool)
    result := []int{}
    for _, item := range A {
        set[item] = true
    }
	for _, item := range B {
		if set[item] {
			result = append(result, item)
		}
	}
	for _, item := range result {
        fmt.Printf("%v\n", item)
	}
	
}