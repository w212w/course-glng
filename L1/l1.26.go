package main

import (
	"fmt"
	"strings"
)

func UniqueChars(s string) bool {
	s = strings.ToLower(s)
	seen := make(map[rune]bool)
	for _, ch := range s {
		if seen[ch] {
			return false
		}
		seen[ch] = true
	}
	return true
}

func main() {
	strings := []string{"abcd", "abCdefAaf", "aabcd"}
	for _, s := range strings {
		fmt.Printf("%q -> %v\n", s, UniqueChars(s))
	}
}
