package main

func InsertionInt(u []int) []int {

	newL := make([]int, len(u))
	copy(newL, u)

	isLess := func(a, b int) bool {
		if a < b {
			return true
		} else {
			return false
		}
	}

	for i := 1; i < len(newL); i++ {
		for j := i; j > 0 && isLess(newL[j], newL[j-1]); j-- {
			old := newL[j]
			newL[j] = newL[j-1]
			newL[j-1] = old
		}
	}

	return newL
}

var insertionImpl string = `func Insertion{{.TypeName}}(u []{{.Type}}) []{{.Type}} {

	newL := make([]{{.Type}}, len(u))
	copy(newL, u)

	isLess := func(a, b {{.Type}}) bool {
		if a < b {
			return true
		} else {
			return false
		}
	}

	for i := 1; i < len(newL); i++ {
		for j := i; j > 0 && isLess(newL[j], newL[j-1]); j-- {
			old := newL[j]
			newL[j] = newL[j-1]
			newL[j-1] = old
		}
	}

	return newL
}


`

var insertionTestTmpl string = `func TestInsertion{{.TypeName}}(t *testing.T) {
	list := []{{.Type}}{4, 2, 1, 7, 5, 0, 3, 9, 6, 8}
	res := Insertion{{.TypeName}}(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != {{.Type}}(i) {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}


`
