package main

import "testing"

func TestInsertionInt(t *testing.T) {
	list := []int{4, 2, 1, 7, 5, 0, 3, 9, 6, 8}
	res := InsertionInt(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != i {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}
