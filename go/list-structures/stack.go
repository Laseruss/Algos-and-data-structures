package main

type Stack struct {
	head   *Node
	length int
}

func NewStack() Stack {
	return Stack{length: 0}
}

func (s *Stack) Push(item string) {
	node := Node{value: item, next: nil}
	if s.length == 0 {
		s.head = &node
		s.length += 1
		return
	}
	s.length += 1
	node.next = s.head
	s.head = &node
}

func (s *Stack) Pop() string {
	if s.length == 0 {
		return "No items in the stack"
	}
	s.length--
	val := s.head.value
	s.head = s.head.next

	return val
}

func (s *Stack) Peek() string {
	if s.length == 0 {
		return "no head"
	}
	return s.head.value
}
