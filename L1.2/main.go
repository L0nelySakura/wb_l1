package main

import (
	"fmt"
	"sync"
)


func main() {

	arr := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup


	for _, number := range arr {
		wg.Add(1)

		go func (num int) {
			defer wg.Done()
			result := num * num
			fmt.Printf("%d^2 = %d\n", num, result)
		}(number)


	}
	wg.Wait()
}