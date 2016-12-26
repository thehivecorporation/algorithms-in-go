package main

import "testing"

func TestQuickInt(t *testing.T) {
	list := []int{4, 2, 1, 3, 0, 5}
	list2 := make([]int, len(list))
	copy(list2, list)

	res := QuickInt(list)

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
		if res[i] != int(i) {
			t.Fatalf("Quicksort algorithm isn't correct: %v\n", res)
		}
	}
}
