package golomb

func golomb(n int, m map[int]int) int {
	if n == 1 {
		return 1
	}

	if _, ok := m[n]; !ok {
		m[n] = 1 + golomb(n-golomb(golomb(n-1, m), m), m)
	}

	return m[n]
}
