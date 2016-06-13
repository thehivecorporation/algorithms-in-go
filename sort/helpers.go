package sort

func isLess(a, b int) bool {
	if a < b {
		return true
	} else {
		return false
	}
}

func swap(u []int, a, b int) {
	old := u[a]
	u[a] = u[b]
	u[b] = old
}