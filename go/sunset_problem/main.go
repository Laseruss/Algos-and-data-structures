// Solving the sunset problem of displaying the buildings that can see the
// sunset. The sun sets on the right side of the slice.
// Using a decreasing monotonic stack.

package main

import (
	"fmt"
)

func main() {
	buildings := []int{18, 14, 13, 16, 12}

	stack := []int{}

	for i, v := range buildings {
		for len(stack) != 0 && v >= buildings[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	fmt.Printf("The building indexes that can see the sunset is: %v\n", stack)
}
