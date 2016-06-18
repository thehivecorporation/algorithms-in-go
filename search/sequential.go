package search

func SequentialInt(u []int, searched int) bool {
	for _, v := range u {
		if v == searched {
			return true
		}
	}

	return false
}