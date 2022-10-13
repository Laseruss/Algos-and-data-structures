package main

import "fmt"

func main() {
	l := []int{1, 5, 3, 350, 23, 210, 535, 22, 23, 48}
	fmt.Println(l)
	fmt.Println(insertionSort(l))
}

func insertionSort(l []int) []int {
	for i := 1; i < len(l); i++ {
		curr := l[i]
		empty := i
		for j := i - 1; j >= 0; j-- {
			if l[j] > curr {
				l[j+1] = l[j]
				empty = j
			}
		}
		l[empty] = curr
	}
	return l
}
