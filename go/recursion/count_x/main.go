package main

import "fmt"

func main() {
	str := "xabcxdhetx"
	fmt.Println(countX(str))
}

func countX(str string) int {
	if len(str) == 1 {
		if str == "x" {
			return 1
		}
		return 0
	}

	if str[0] == 'x' {
		return 1 + countX(str[1:])
	}
	return countX(str[1:])
}
