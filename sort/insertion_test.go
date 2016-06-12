package sort

import (
	"testing"
	"fmt"
)

func TestInsertion(t *testing.T) {
	list := []int{4, 2, 1, 7, 5, 0, 3, 9, 6, 8}
	res := InsertionInt(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d", len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != i {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}

func TestUnderlying(t *testing.T) {
	list := make([]int,2)

	list[0] = 5
	list[1] = 9

	res := testUnderlying(list)

	fmt.Printf("%v\n %v\n", res, list)
}