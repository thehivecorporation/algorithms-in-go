package main

import "testing"

func TestSelectionUint8(t *testing.T) {
	list := []uint8{4, 2, 1, 7, 3, 0, 5, 6}
	res := SelectionUint8(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint8(i) {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}
