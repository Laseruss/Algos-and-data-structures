package doublylinked

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	val  any
	prev *Node
	next *Node
}

func MakeQueue(val any) Queue {
	node := Node{val, nil, nil}
	return Queue{
		head:   &node,
		tail:   &node,
		length: 1,
	}
}

func (d *Queue) Enque(val any) {
	d.length++
	node := Node{
		val:  val,
		prev: d.tail,
		next: nil,
	}

	if d.length == 0 {
		d.head = &node
		d.tail = &node
		return
	}

	d.tail.next = &node
	d.tail = &node
}

func (d *Queue) Deque() any {
	if d.tail == nil {
		return "nothing in the queue"
	}

	val := d.head.val
	d.length--

	if d.length == 0 {
		d.head = nil
		d.tail = nil

		return val
	}

	d.head = d.head.next
	d.head.prev = nil
	return val
}

func (d *Queue) Peek() any {
	if d.length == 0 {
		return "nothing in the queue"
	}

	return d.head.val
}
