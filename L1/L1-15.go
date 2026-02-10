package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(n int) string {
	return strings.Repeat("x", n)
}

func someFunc() {
	v := createHugeString(1 << 20)
	justString = strings.Clone(v[:100])
	fmt.Println("fixed:", len(justString))
}

func main() {
	someFunc()
}
