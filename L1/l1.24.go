package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {

	x1, y1 := 3.0, 4.0
	x2, y2 := 7.0, 1.0

	p1 := NewPoint(x1, y1)
	p2 := NewPoint(x2, y2)

	distance := p1.Distance(p2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
