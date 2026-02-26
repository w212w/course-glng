package main

import "strings"

func reverseSentence(s string) string {
	spl := strings.Split(s, " ")
	for i, j := 0, len(spl)-1; i < j; i, j = i+1, j-1 {
		spl[i], spl[j] = spl[j], spl[i]
	}
	return strings.Join(spl, " ")
}

func revRunes(r []rune, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}

func reverseWordsInSentence(s string) string {
	r := []rune(s)
	n := len(r)
	if n == 0 {
		return s
	}
	revRunes(r, 0, n-1)
	start := 0
	for i := 0; i <= n; i++ {
		if i == n || r[i] == ' ' {
			revRunes(r, start, i-1)
			start = i + 1
		}
	}
	return string(r)
}

func main() {
	println(reverseSentence("Hello, World! Привет, мир!"))
	println(reverseWordsInSentence("Hello, World! Привет, мир!"))
}
