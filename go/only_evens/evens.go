package onlyevens

func onlyEvens(l []int, new []int) []int {
	if len(l) == 0 {
		return new
	}

	if l[0]%2 == 0 {
		new = append(new, l[0])
	}

	return onlyEvens(l[1:], new)
}
