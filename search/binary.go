package search

func BinaryInt(u []int, val int, sortingAlgorithm func(v []int) []int) bool {
	newL := sortingAlgorithm(u)
	return binaryInt(newL, val)
}

func binaryInt(u []int, val int) bool {
	low := 0
	high := len(u) - 1

	for low <= high {
		ix := (low + high) / 2
		if val == u[ix] {
			return true
		} else if val < u[ix] {
			high = ix - 1
		} else {
			low = ix + 1
		}
	}

	return false
}
