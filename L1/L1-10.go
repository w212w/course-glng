package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	dict := make(map[int][]float64)
	for _, num := range nums {
		key := int(num/10) * 10
		dict[key] = append(dict[key], num)
	}

	keys := make([]int, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d:{", k)
		for i, v := range dict[k] {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%.1f", v)
		}
		fmt.Println("}")
	}

}
