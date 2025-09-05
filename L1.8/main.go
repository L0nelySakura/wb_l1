package main

import (
	"fmt"
	"os"
	"strconv"
	"math"
)

// go run .\main.go 'number' 'position' 'value'
func main() {
	var a int
	number, _ := strconv.Atoi(os.Args[1])
	position, _ := strconv.Atoi(os.Args[2])
	value, _ := strconv.Atoi(os.Args[3])

	if value == 0 {
    	a = (number &^ int(math.Pow(2, float64(position - 1))))

	} else {
    	a = (number | int(math.Pow(2, float64(position - 1))))
    	
	}
	fmt.Println(a)
}