package main

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

var quickImpl string = `func Quick{{.TypeName}}(u []{{.Type}}) []{{.Type}} {
	length := len(u)

	if length <= 1 {
		sliceCopy := make([]{{.Type}}, length)
		copy(sliceCopy, u)
		return sliceCopy
	}

	m := u[rand.Intn(length)]

	less := make([]{{.Type}}, 0, length)
	middle := make([]{{.Type}}, 0, length)
	more := make([]{{.Type}}, 0, length)

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

	less, more = Quick{{.TypeName}}(less), Quick{{.TypeName}}(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}


`

var quickTest string = `func TestQuick{{.TypeName}}(t *testing.T) {
	list := []{{.Type}}{4, 2, 1, 3, 0, 5}
	list2 := make([]{{.Type}}, len(list))
	copy(list2, list)

	res := Quick{{.TypeName}}(list)

	//Check that we haven't modified the original list
	for k, v := range list {
		if list2[k] != v {
			t.Error("Original list was modified")
		}
	}

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != {{.Type}}(i) {
			t.Fatalf("Quicksort algorithm isn't correct: %v\n", res)
		}
	}
}

`
