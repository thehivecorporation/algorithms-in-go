package sort

import (
	"math/rand"
)

func QuickInt(u []int) []int {
	length := len(u)

	if length <= 1 {
		sliceCopy := make([]int, length)
		copy(sliceCopy, u)
		return sliceCopy
	}

	m := u[rand.Intn(length)]

	less := make([]int, 0, length)
	middle := make([]int, 0, length)
	more := make([]int, 0, length)

	for _, item := range u {
		switch {
		case item < m:
			less = append(less, item)
		case item == m:
			middle = append(middle, item)
		case item > m:
			more = append(more, item)
		}
	}

	less, more = QuickInt(less), QuickInt(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}
