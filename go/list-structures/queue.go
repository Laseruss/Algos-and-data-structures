package main

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

func NewQueue() Queue {
	return Queue{head: nil, tail: nil, length: 0}
}

func (q *Queue) Enque(item string) {
	node := Node{value: item}

	if q.length == 0 {
		q.length++
		q.head = &node
		q.tail = &node
		return
	}

	q.length++
	q.tail.next = &node
	q.tail = &node
}

func (q *Queue) Deque() string {
	if q.length == 0 {
		return "No items in the queue"
	}

	q.length--
	val := q.head.value
	q.head = q.head.next

	return val
}

func (q *Queue) Peek() string {
	if q.length == 0 {
		return "No items in the queue"
	}

	return q.head.value
}
