package search

import (
	"testing"

	"github.com/sayden/algorithms-in-go/sort"
)

func TestBinaryInt(t *testing.T) {
	list := []int{4, 2, 1, 7, 3, 0, 5, 6}

	res := BinaryInt(list, 2, sort.BubbleInt)
	if !res {
		t.Error("Value not found when it should")
	}

	res = BinaryInt(list, 33, sort.HeapInt)
	if res {
		t.Error("Value found when it shouldn't")
	}
}
