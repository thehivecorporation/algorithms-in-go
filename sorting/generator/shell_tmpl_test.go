package main

import "testing"

func TestShellUint8(t *testing.T) {
	list := []uint8{4, 2, 1, 3, 0, 9, 6, 8, 7, 5}
	res := ShellUint8(list, 0)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint8(i) {
			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
		}
	}
}
