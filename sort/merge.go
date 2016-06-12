package sort

func MergeInt(u []int) (newL []int) {

	if len(u) <= 1 {
		return u
	}

	middlePoint := len(u) / 2

	leftHalf := u[:middlePoint]
	rightHalf := u[middlePoint:]

	left := MergeInt(leftHalf)
	right := MergeInt(rightHalf)

	merge := func(left, right []int) []int {
		var result []int
		for len(left) > 0 || len(right) > 0 {
			if len(left) > 0 && len(right) > 0 {
				if left[0] <= right[0] {
					result = append(result, left[0])
					left = left[1:]
				} else {
					result = append(result, right[0])
					right = right[1:]
				}
			} else if len(left) > 0 {
				result = append(result, left[0])
				left = left[1:]
			} else if len(right) > 0 {
				result = append(result, right[0])
				right = right[1:]
			}
		}

		return result
	}

	return merge(left, right)
}
