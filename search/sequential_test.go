package search

import "testing"

func TestSequentialInt(t *testing.T) {
	list := []int{4, 2, 1, 7, 3, 0, 5, 6}
	res := SequentialInt(list, 3)
	if !res {
		t.Error("Value that existed wasn't found")
	}

	res = SequentialInt(list, 33)
	if res {
		t.Error("Values that didn't exist was found")
	}
}
