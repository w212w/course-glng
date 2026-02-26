package main

import "fmt"

func main() {
	a, b := 12, 34
	fmt.Printf("Исходные: a=%d, b=%d\n", a, b)

	x, y := a, b
	x = x + y // 46
	y = x - y // 12
	x = x - y // 34

	fmt.Printf("После обмена (сложение/вычитание): a=%d, b=%d\n", x, y)
}
