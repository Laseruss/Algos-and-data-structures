package main

import "fmt"

func main() {
	haystack := []int{3, 5, 9, 10, 34}
	needle := 9
	fmt.Println(haystack, needle, linearSearch(haystack, needle))
}

func linearSearch(l []int, needle int) int {
	for i := range l {
		if l[i] == needle {
			return i
		}
	}
	return -1
}
