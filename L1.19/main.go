package main

import (
	"fmt"
)

func reverseString(input string) string{
	runes := []rune(input)
	for i := 0; i < (len(runes)) / 2; i++ {
		runes[i], runes[len(runes) - 1 - i] = runes[len(runes) - 1 - i], runes[i]
	}
	return string(runes)

}


func main() {
	words := []string{
		"hello",
		"привет",
		"1234567",
		"123456",
		"12",
		"1",
		"",
	}

	for _, word := range words {
		fmt.Printf("%s\t%s\n", word, reverseString(word))
	}

}