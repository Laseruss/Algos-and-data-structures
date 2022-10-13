package main

import "fmt"

func main() {
	l := []int{5, 3, 2, 9, 7}
	doubleArray(l, 0)
	fmt.Println(l)
}

func doubleArray(l []int, index int) {
	l[index] *= 2
	if index == len(l)-1 {
		return
	}
	index++
	doubleArray(l, index)
}
