package main

func SelectionUint8(u []uint8) []uint8 {
	newList := make([]uint8, len(u))
	tempList := make([]uint8, len(u))
	copy(tempList, u)

	for i := 0; len(tempList) > 0; i++ {
		min := tempList[0]
		index := 0

		for k := 0; k < len(tempList); k++ {
			if tempList[k] < min {
				index = k
				min = tempList[k]
			}
		}

		newList[i] = min
		tempList = append(tempList[:index], tempList[index+1:]...)
	}

	return newList
}

var selectionImpl string = `func Selection{{.TypeName}}(u []{{.Type}}) []{{.Type}} {
	newList := make([]{{.Type}}, len(u))
	tempList := make([]{{.Type}}, len(u))
	copy(tempList, u)

	for i := 0; len(tempList) > 0; i++ {
		min := tempList[0]
		index := 0

		for k := 0; k < len(tempList); k++ {
			if tempList[k] < min {
				index = k
				min = tempList[k]
			}
		}

		newList[i] = min
		tempList = append(tempList[:index], tempList[index+1:]...)
	}

	return newList
}

`

var selectionTest string = `func TestSelection{{.TypeName}}(t *testing.T) {
	list := []{{.Type}}{4, 2, 1, 7, 3, 0, 5, 6}
	res := Selection{{.TypeName}}(list)

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
