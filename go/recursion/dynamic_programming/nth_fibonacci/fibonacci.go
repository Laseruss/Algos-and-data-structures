package fibonacci

func nthFibonacci(n uint, m map[uint]uint) uint {
	if n == 0 || n == 1 {
		return n
	}

	if _, ok := m[n]; !ok {
		m[n] = nthFibonacci(n-2, m) + nthFibonacci(n-1, m)
	}

	return m[n]
}

func bottomUpFibonacci(n uint) uint {
	if n == 0 {
		return 0
	}
	var a, b uint = 0, 1

	for i := uint(1); i < n; i++ {
		a, b = b, a+b
	}

	return b
}
