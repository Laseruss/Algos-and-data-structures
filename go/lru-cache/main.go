package main

import "fmt"

type LRUCache struct {
	Cap    int
	Values map[int]*Node
	Least  *List
}

type List struct {
	Len  int
	Head *Node
	Tail *Node
}

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

func (l *List) AddNode(key, val int) {
	newNode := &Node{
		Key:  key,
		Val:  val,
		Next: nil,
		Prev: nil,
	}

	l.Len++
	if l.Len == 1 {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	l.Head.Prev = newNode
	newNode.Next = l.Head
	l.Head = newNode
}

func (l *List) SwapNode(node *Node) {
	if l.Len == 1 {
		return
	} else if node == l.Head {
		return
	} else if node == l.Tail {
		l.Tail = node.Prev
		node.Prev = nil

		node.Next = l.Head
		l.Head.Prev = node
		l.Head = node
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		node.Next = l.Head
		l.Head.Prev = node
		node.Prev = nil
		l.Head = node
	}
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:    capacity,
		Values: make(map[int]*Node),
		Least: &List{
			Len:  0,
			Head: nil,
			Tail: nil,
		},
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Values[key]; ok {
		this.Least.SwapNode(node)
		fmt.Println(node.Val)
		return node.Val
	}
	fmt.Println(-1)
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// If the key already in the cache, move it up to the most recently used.
	// Update the value for the given key
	if node, ok := this.Values[key]; ok {
		node.Val = value
		this.Least.SwapNode(node)
		return
	}

	// Else add one to the length set the newly added key, value as the head of the list
	this.Least.AddNode(key, value)
	this.Values[key] = this.Least.Head

	if this.Least.Len > this.Cap {
		delete(this.Values, this.Least.Tail.Key)
		this.Least.Tail = this.Least.Tail.Prev
		fmt.Println("shrinking: ", this.Least.Tail)
		this.Least.Len--
	}
}

func main() {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3)
	cache.Put(4, 1)
	cache.Get(1)
	cache.Get(2)

	fmt.Println(cache.Values)
}
