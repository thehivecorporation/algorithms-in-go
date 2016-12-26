package main

func SequentialInt(u []int, searched int) bool {
	for _, v := range u {
		if v == searched {
			return true
		}
	}

	return false
}

var sequentialImpl string = `func Sequential{{.TypeName}}(u []{{.Type}}, searched {{.Type}}) bool {
	for _, v := range u {
		if v == searched {
			return true
		}
	}

	return false
}

`

var sequentialTest string = `func TestSequential{{.TypeName}}(t *testing.T) {
	list := []{{.Type}}{4, 2, 1, 7, 3, 0, 5, 6}
	res := Sequential{{.TypeName}}(list, 3)
	if !res {
		t.Error("Value that existed wasn't found")
	}

	res = Sequential{{.TypeName}}(list, 33)
	if res {
		t.Error("Values that didn't exist was found")
	}
}


`
