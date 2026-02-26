package main

import "fmt"

func main() {
	array := []string{"cat", "cat", "dog", "cat", "tree", "tree", "cat"}
	set := make(map[string]struct{})
	for _, v := range array {
		set[v] = struct{}{}
	}
	unique := make([]string, 0, len(set))
	for k, _ := range set {
		unique = append(unique, k)
	}
	fmt.Println("Множество =", unique)
}
