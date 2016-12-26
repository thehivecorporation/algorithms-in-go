package main

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

func heapInt(u []int) []int {
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

var heapTmpl string = `func heapify{{.TypeName}}(u []{{.Type}}, i, max int) {
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
		common.Swap{{.TypeName}}(u, i, largest)

		heapify{{.TypeName}}(u, largest, max)
	}
}

func Heap{{.TypeName}}(u []{{.Type}}) []{{.Type}} {
	newL := make([]{{.Type}}, len(u))
	copy(newL, u)

	for i := (len(u) / 2) - 1; i > 0; i-- {
		heapify{{.TypeName}}(newL, i, len(u))
	}

	for i := len(u) - 1; i > 0; i-- {
		common.Swap{{.TypeName}}(newL, 0, i)

		heapify{{.TypeName}}(newL, 0, i)
	}

	return newL
}


`

var heapTestTmpl string = `
func TestHeap{{.TypeName}}(t *testing.T) {
	list := []{{.Type}}{4, 2, 1, 3, 0}
	res := Heap{{.TypeName}}(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != {{.Type}}(i) {
			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
		}
	}
}
`
