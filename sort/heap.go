package sort

func HeapInt(u []int) (newL []int) {
	newL = make([]int, len(u))
	copy(newL, u)

	siftDown := func(u []int, lo, hi, first int) {
		root := lo

		for {
			child := 2*root + 1
			if child >= hi {
				break
			}
			if child+1 < hi && isLess(u[first+child], u[first+child+1]) {
				child++
			}
			if !isLess(u[first+root], u[first+child]) {
				return
			}

			swap(u, first+root, first+child)

			root = child
		}
	}

	first := 0
	lo := 0
	hi := len(newL)

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(newL, i, hi, first)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		swap(u, first, first+1)
		siftDown(newL, lo, i, first)
	}

	return newL
}
