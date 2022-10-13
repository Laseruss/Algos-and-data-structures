package main

import "fmt"

func main() {
	heap := Heap{20, 15, 17, 12, 10, 9}

	fmt.Println(heap)

	for range heap {
		heap = heap.Delete()
		fmt.Println(heap)
	}
}
