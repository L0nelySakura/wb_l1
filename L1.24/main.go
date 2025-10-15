package main


import (
    "fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) Point {
	return Point{
		x: x,
		y: y,
	}

}

func (p Point) Distance(other Point) float64 {
	return math.Sqrt(math.Pow(p.x - other.x, 2) + math.Pow(p.y - other.y, 2))
}


func main() {
	point1 := NewPoint(1.5, 2.0)
    point2 := NewPoint(4.0, 6.0)
    fmt.Printf("Расстояние между точками: %f\n", point2.Distance(point1))
}
