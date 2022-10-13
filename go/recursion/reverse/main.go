package main

import "fmt"

func main() {
	str := "abcde"
	fmt.Println(reverse(str))
}

func reverse(str string) string {
	if len(str) == 1 {
		return str
	}

	return reverse(str[1:]) + string(str[0])
}
