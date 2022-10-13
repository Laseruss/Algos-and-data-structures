package add_until

func addUntilX(l []int, max int) int {
	if len(l) == 0 {
		return 0
	}

	sumOfRest := addUntilX(l[1:], max)

	if l[0]+sumOfRest > max {
		return sumOfRest
	} else {
		return l[0] + sumOfRest
	}
}
