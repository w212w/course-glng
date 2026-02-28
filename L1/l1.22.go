package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b int = 4_500_000, 2_300_000

	fmt.Println("=== Операции с int ===")
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", b)
	fmt.Printf("a + b = %d\n", a+b)
	fmt.Printf("a - b = %d\n", a-b)
	fmt.Printf("a * b = %d\n", a*b)
	fmt.Printf("a / b = %d\n", a/b)

	fmt.Println("=== Операции с math/big ===")

	bigA := new(big.Int).SetInt64(4_500_000_000_000_000_000)
	bigB := new(big.Int).SetInt64(2_300_000_000_000_000_000)

	fmt.Printf("a = %s\n", bigA)
	fmt.Printf("b = %s\n", bigB)

	sum := new(big.Int).Add(bigA, bigB)
	fmt.Printf("a + b = %s\n", sum)

	diff := new(big.Int).Sub(bigA, bigB)
	fmt.Printf("a - b = %s\n", diff)

	product := new(big.Int).Mul(bigA, bigB)
	fmt.Printf("a * b = %s\n", product)

	quotient := new(big.Int).Div(bigA, bigB)
	fmt.Printf("a / b = %s\n", quotient)
}
