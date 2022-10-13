package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

func main() {
	queue := NewQueue()
	words := []string{"hello", "world", "my", "name", "is", "Micke"}

	for _, w := range words {
		queue.Enque(w)
	}

	for queue.length > 0 {
		val := queue.Deque()
		fmt.Println(val)
	}

	stack := NewStack()
	words = []string{"hello", "world", "my", "name", "is", "Micke"}

	for _, w := range words {
		stack.Push(w)
	}

	for stack.length > 0 {
		val := stack.Pop()
		fmt.Println(val)
	}
}
