package main

import (
	"fmt"
	"math"
	"time"
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

// Собственная функция Sleep через busy-wait
func Sleep(duration time.Duration) {
	start := time.Now()
	for time.Since(start) < duration {
		// Пустой цикл ожидания
	}
}

func main() {
	fmt.Println("Ожидание 2 секунды...")
	Sleep(2 * time.Second)
	fmt.Println("Прошло 2 секунды!")
}
