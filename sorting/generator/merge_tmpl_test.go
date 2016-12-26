package main

import "testing"

func TestMergeInt(t *testing.T) {
	list := []int{4, 2, 1, 3, 0, 9, 6, 8, 7, 5, 10}
	res := MergeInt(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int(i) {
			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
		}
	}
}
