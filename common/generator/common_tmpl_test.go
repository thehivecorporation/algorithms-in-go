package main

import "testing"

func TestSwapInt(t *testing.T) {
	ar := []int{1, 2}

	swapInt(ar, 0, 1)
	if ar[0] != 2 || ar[1] != 1 {
		t.Fail()
	}
}

func TestIsLessInt(t *testing.T) {
	if !isLessInt(1, 2) {
		t.Fail()
	}

	if isLessInt(2, 1) {
		t.Fail()
	}
}
