package main

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

var binaryImpl string = `func Binary{{.TypeName}}(u []{{.Type}}, val {{.Type}}, sortingAlgorithm func(v []{{.Type}}) []{{.Type}}) bool {
	newL := sortingAlgorithm(u)
	return binary{{.TypeName}}(newL, val)
}

func binary{{.TypeName}}(u []{{.Type}}, val {{.Type}}) bool {
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

`

var binaryTest string = `func TestBinary{{.TypeName}}(t *testing.T) {
	list := []{{.Type}}{4, 2, 1, 7, 3, 0, 5, 6}

	res := Binary{{.TypeName}}(list, 2, sorting.Bubble{{.TypeName}})
	if !res {
		t.Error("Value not found when it should")
	}

	res = Binary{{.TypeName}}(list, 33, sorting.Heap{{.TypeName}})
	if res {
		t.Error("Value found when it shouldn't")
	}
}

`
