package main


import (
    "fmt"
)


func remove(slice []int, i int) []int {
	if i < 0 || len(slice) <= i {
		return slice
	}

	copy(slice[i:], slice[i+1:])
	slice[len(slice) - 1] = 0
	return slice[:len(slice) - 1]

}


func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	new_slice := remove(slice, 2)
	fmt.Println(new_slice)
}
