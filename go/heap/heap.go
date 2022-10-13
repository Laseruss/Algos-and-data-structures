package main

type Heap []int

func (h Heap) Insert(val int) Heap {
	h = append(h, val)
	insertedInd := len(h) - 1
	parentInd := parent(insertedInd)
	for insertedInd < len(h) && h[insertedInd] > h[parentInd] {
		h[insertedInd], h[parentInd] = h[parentInd], h[insertedInd]
		insertedInd = parentInd
		parentInd = parent(insertedInd)
	}

	return h
}

func (h Heap) Delete() Heap {
	h[0] = h[len(h)-1]

	ind := 0

	for left(ind) < len(h) && right(ind) < len(h) {
		var swapInd int
		if h[left(ind)] > h[right(ind)] {
			swapInd = left(ind)
		} else {
			swapInd = right(ind)
		}

		if h[swapInd] > h[ind] {
			h[swapInd], h[ind] = h[ind], h[swapInd]
			ind = swapInd
			continue
		}
		break
	}

	return h[:len(h)-1]
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return (index * 2) + 1
}

func right(index int) int {
	return (index * 2) + 2
}
