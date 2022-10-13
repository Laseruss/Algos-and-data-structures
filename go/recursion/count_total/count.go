package count

func countTotal(strs []string) int {
	if len(strs) == 0 {
		return 0
	}

	return len(strs[0]) + countTotal(strs[1:])
}
