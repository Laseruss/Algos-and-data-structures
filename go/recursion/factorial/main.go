package main

import "fmt"

func main() {
	num := 6
	fmt.Println(factorial(num))
}

func factorial(x int) int {
	if x == 1 {
		return 1
	}

	return x * factorial(x-1)
}
