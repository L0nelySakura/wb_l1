package main

import (
	"fmt"
)

func main() {
    sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]bool)
    result := []string{}

    for _, item := range sequence {
        set[item] = true
    }
	for key, _ := range set {
		result = append(result, key)
	}
	for _, item := range result {
		fmt.Printf("%s\n", item)
	}

}