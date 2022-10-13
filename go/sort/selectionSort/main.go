package main

import "fmt"

func main() {
	l := []int{1, 5, 3, 350, 23, 210, 535, 22, 23, 48}
	fmt.Println(l)
	fmt.Println(selectionSort(l))
}

func selectionSort(l []int) []int {
	for i := range l {
		smallestIndex := i

		for j := i + 1; j < len(l); j++ {
			if l[smallestIndex] > l[j] {
				smallestIndex = j
			}
		}
		l[i], l[smallestIndex] = l[smallestIndex], l[i]
	}
	return l
}
