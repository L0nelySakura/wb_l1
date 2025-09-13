package main

import (
	"fmt"
)

func main() {	
	numbers := []int{1, 2, 3, 5}

	inChan := make(chan int)
	outChan := make(chan int)

	go func() {
		for _, x := range numbers {
			inChan <- x
		}
		close(inChan)
	}()

	go func() {
		for x := range inChan {
			outChan <- x * 2
		}
		close(outChan)
	}()

	for result := range outChan {
		fmt.Println(result)
	}

}