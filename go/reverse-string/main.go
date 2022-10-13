package main

import "fmt"

func main() {
	s := "hello world"
	fmt.Println(reverse(s))
}

func reverse(s string) string {
	runes := []rune(s)
	l := 0
	r := len(runes) - 1

	for l < r {
		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}
	return string(runes)
}
