package sort

import "github.com/sayden/algorithms-in-go/common"

func heapifyInt(u []int, i, max int) {
	left := 2 * i
	right := left + 1
	largest := 0

	if left < max && u[left] > u[i] {
		largest = left
	} else {
		largest = i
	}

	if right < max && u[right] > u[largest] {
		largest = right
	}
	if largest != i {
		common.SwapInt(u, i, largest)

		heapifyInt(u, largest, max)
	}
}

func HeapInt(u []int) []int {
	newL := make([]int, len(u))
	copy(newL, u)

	for i := (len(u) / 2) - 1; i > 0; i-- {
		heapifyInt(newL, i, len(u))
	}

	for i := len(u) - 1; i > 0; i-- {
		common.SwapInt(newL, 0, i)

		heapifyInt(newL, 0, i)
	}

	return newL
}
