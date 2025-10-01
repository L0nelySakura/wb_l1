package main

import (
	"fmt"
	"strings"
)

func reverseString(input string) string{
	words := strings.Fields(input)
	for i := 0; i < (len(words)) / 2; i++ {
		words[i], words[len(words) - 1 - i] = words[len(words) - 1 - i], words[i]
	}
	return strings.Join(words, " ")

}


func main() {
	words := []string{
		"sun dog snow",
		"sun dog",
		"sun",
		"",
	}

	for _, word := range words {
		fmt.Printf("%s\t->%s\n", word, reverseString(word))
	}

}