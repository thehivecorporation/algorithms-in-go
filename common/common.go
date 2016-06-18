package common

func SwapInt(u []int, a, b int) {
	old := u[b]
	u[b] = u[a]
	u[a] = old
}

func isLessInt(a, b int) bool {
	if a < b {
		return true
	} else {
		return false
	}
}