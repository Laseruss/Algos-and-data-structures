package main

import "fmt"

func main() {
	steps := 5
	fmt.Println(numberOfPaths(steps))
}

func numberOfPaths(n int) int {
	if n < 0 {
		return 0
	}
	if n == 1 || n == 0 {
		return 1
	}
	return numberOfPaths(n-1) + numberOfPaths(n-2) + numberOfPaths(n-3)
}
