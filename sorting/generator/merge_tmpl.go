package main

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

var mergeImpl string = `func Merge{{.TypeName}}(u []{{.Type}}) (newL []{{.Type}}) {

	if len(u) <= 1 {
		return u
	}

	middlePoint := len(u) / 2

	leftHalf := u[:middlePoint]
	rightHalf := u[middlePoint:]

	left := Merge{{.TypeName}}(leftHalf)
	right := Merge{{.TypeName}}(rightHalf)

	merge := func(left, right []{{.Type}}) []{{.Type}} {
		var result []{{.Type}}
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


`

var mergeTest string = `
func TestMerge{{.TypeName}}(t *testing.T) {
	list := []{{.Type}}{4, 2, 1, 3, 0, 9, 6, 8, 7, 5, 10}
	res := Merge{{.TypeName}}(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != {{.Type}}(i) {
			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
		}
	}
}
`
