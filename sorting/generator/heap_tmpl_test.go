package main

import "testing"

func TestHeapInt(t *testing.T) {
	list := []int{4, 2, 1, 3, 0}
	res := heapInt(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int(i) {
			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
		}
	}
}
