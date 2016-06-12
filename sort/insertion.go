package sort

func InsertionInt(l []int) []int {
	for i := 1; i < len(l); i++ {
		val := l[i]

		for k := i - 1; k >= 0; k-- {
			prev := l[k]

			if val < prev {
				l[k+1] = prev
				l[k] = val
			}
		}
	}

	return l
}

func testUnderlying(l []int) []int {
	res := l

	res[0] = 10

	return res
}