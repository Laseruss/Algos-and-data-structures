package main

import (
	"fmt"
)

func main() {
	haystack := []int{1, 4, 9, 15, 53, 100, 342, 1539}
	fmt.Println(haystack)
	for _, v := range haystack {
		needle := v
		fmt.Printf("The needle is at index %d\n", binarySearch(haystack, needle))
	}
}

func binarySearch(haystack []int, needle int) int {
	low, high := 0, len(haystack)-1
	for low <= high {
		index := (high + low) / 2
		if haystack[index] == needle {
			return index
		} else if needle < haystack[index] {
			high = index
		} else {
			low = index + 1
		}
	}
	return -1
}
