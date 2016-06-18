package common

func SwapInt(u []int, a, b int) {
	old := u[b]
	u[b] = u[a]
	u[a] = old
}
