package main

import (
	"fmt"
	"math"
)

func main() {
	// Довольно странная задача, как классифицировать значения от -9.9 до 9.9?
	// Сделал так, что значения от -10 не включительно до 0 не включительно попадают в -0
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[string][]float64)
	var key string

	for _, t := range temps {
		if t >= 0 {
			key = fmt.Sprintf("%d", int(math.Floor(t / 10)) * 10)
		} else {
			key = fmt.Sprintf("-%d", int(math.Floor(-t / 10)) * 10)
		}
		groups[key] = append(groups[key], t)
	}

	for key, values := range groups {
        fmt.Printf("%s: %v\n", key, values)
    }
}