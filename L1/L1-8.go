package main

import (
	"fmt"
)

func setBit(n int64, i uint, bit uint) int64 {
	if bit == 1 {
		return n | (1 << i)
	}
	return n &^ (1 << i)
}

func main() {
	var n int64 = 5
	fmt.Printf("Исходное число: %d (%04b)\n", n, n)

	n1 := setBit(n, 1, 0)
	fmt.Printf("Установка 1-го бита в 0: %d (%04b)\n", n1, n1)

	n2 := setBit(n, 2, 1)
	fmt.Printf("Установка 2-го бита в 1: %d (%04b)\n", n2, n2)

	n3 := setBit(n, 3, 1)
	fmt.Printf("Установка 3-го бита в 1: %d (%04b)\n", n3, n3)
}
