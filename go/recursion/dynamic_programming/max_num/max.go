package max

func maxNum(l []int) int {
	if len(l) == 1 {
		return l[0]
	}
	maxOfRest := maxNum(l[1:])

	if l[0] > maxOfRest {
		return l[0]
	} else {
		return maxOfRest
	}
}
