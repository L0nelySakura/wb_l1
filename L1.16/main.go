package main


import (
	"fmt"
)


func quickSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	mid := list[len(list) / 2]
	var left, middle, right []int

	for _, number := range list {
		switch {
		case number < mid:
			left = append(left, number)
		case number == mid:
			middle = append(middle, number)
		case number > mid:
			right = append(right, number)
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, middle...), right...)
}



func main() {
	fmt.Println(quickSort([]int{2345, 641, 63493, 1247, 563496, 1432, 53, 62}))
}