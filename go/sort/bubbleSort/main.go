package main

import "fmt"

func main() {
	l := []int{1, 5, 3, 350, 23, 210, 535, 22, 23, 48}
	fmt.Println(l)
	bubbleSort(l)
	fmt.Println(l)
}

func bubbleSort(l []int) {
	for i := range l {
		for j := range l[:len(l)-1-i] {
			if l[j] > l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
			}
		}
	}
}
