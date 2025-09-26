package main

import (
	"fmt"
)


func main() {
	arr := []int{1, 2, 5, 7, 9, 10, 100}
	goal := 10
	
	
	left, right, curr := 0, len(arr) - 1, 0
	loop:
		for right - left > 1 {
			curr = (right - left) / 2 + left
			switch {
			case arr[curr] < goal:
				left = curr
			case arr[curr] == goal:
				break loop
			case arr[curr] > goal:
				right = curr
			}
		}
		
	if arr[curr] == goal {
		fmt.Println(curr)
	} else {
		fmt.Println(-1)
	}
}