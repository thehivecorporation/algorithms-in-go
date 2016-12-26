package main

func bubbleInt(u []int) (newList []int) {
	newList = make([]int, len(u))
	copy(newList, u)

	for {
		hasSwap := false
		for i := 0; i < len(newList)-1; i++ {
			if newList[i] > newList[i+1] {
				old := newList[i]
				newList[i] = newList[i+1]
				newList[i+1] = old
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}

	return newList
}

var bubbleNumericTmpl string = `
func Bubble{{.TypeName}}(u []{{.Type}}) (newList []{{.Type}}) {
	newList = make([]{{.Type}}, len(u))
	copy(newList, u)

	for {
		hasSwap := false
		for i := 0; i < len(newList)-1; i++ {
			if newList[i] > newList[i+1] {
				old := newList[i]
				newList[i] = newList[i+1]
				newList[i+1] = old
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}

	return newList
}

`

var bubbleNumericTestTmpl string = `
func TestBubble{{.TypeName}}(t *testing.T) {
	list := []{{.Type}}{4, 2, 1, 7, 3, 0, 5, 6}
	res := Bubble{{.TypeName}}(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != {{.Type}}(i) {
			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
		}
	}
}

func BenchmarkBubble{{.TypeName}}(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]{{.Type}}, 1000)
	for i := 0; i < 1000; i++ {
		list[i] = {{.Type}}(rand.Int63())
	}

	b.ResetTimer()

	for i:=0; i < b.N; i++{
		Bubble{{.TypeName}}(list)
	}
}

`
