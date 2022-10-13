package doublylinked

import "testing"

func TestQueue(t *testing.T) {
	queue := MakeQueue(5)
	queue.Enque(10)
	queue.Deque()

	if queue.Peek() != 10 || queue.tail.val != 10 || queue.head.val != 10 {
		t.Errorf("did not get what was expected")
	}
}
