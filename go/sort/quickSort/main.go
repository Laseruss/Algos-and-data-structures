package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := []int{10, 5, 2, 3, 323, 56, 45, 8}
	fmt.Println(s)
	fmt.Println(sort(s))
}

// Implementation of quick sort
func sort(l []int) []int {
	if len(l) < 2 {
		return l
	}

	lesser, greater := 0, len(l)-1
	pivot := rand.Int() % len(l)

	l[pivot], l[greater] = l[greater], l[pivot]

	for i := range l {
		if l[i] < l[greater] {
			l[lesser], l[i] = l[i], l[lesser]
			lesser++
		}
	}

	l[lesser], l[greater] = l[greater], l[lesser]

	sort(l[:lesser])
	sort(l[lesser+1:])

	return l
}
