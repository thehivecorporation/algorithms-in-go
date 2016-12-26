// DO NOT EDIT THIS FILE MANUALLY
// Generated by go generate
package sorting

import "testing"

func TestMergeUint(t *testing.T) {
	list := []uint{4, 2, 1, 3, 0, 9, 6, 8, 7, 5, 10}
	res := MergeUint(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint(i) {
			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
		}
	}
}

func TestMergeUint8(t *testing.T) {
	list := []uint8{4, 2, 1, 3, 0, 9, 6, 8, 7, 5, 10}
	res := MergeUint8(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint8(i) {
			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
		}
	}
}

func TestMergeUint16(t *testing.T) {
	list := []uint16{4, 2, 1, 3, 0, 9, 6, 8, 7, 5, 10}
	res := MergeUint16(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint16(i) {
			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
		}
	}
}

func TestMergeUint32(t *testing.T) {
	list := []uint32{4, 2, 1, 3, 0, 9, 6, 8, 7, 5, 10}
	res := MergeUint32(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint32(i) {
			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
		}
	}
}

func TestMergeUint64(t *testing.T) {
	list := []uint64{4, 2, 1, 3, 0, 9, 6, 8, 7, 5, 10}
	res := MergeUint64(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint64(i) {
			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
		}
	}
}

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

func TestMergeInt8(t *testing.T) {
	list := []int8{4, 2, 1, 3, 0, 9, 6, 8, 7, 5, 10}
	res := MergeInt8(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int8(i) {
			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
		}
	}
}

func TestMergeInt16(t *testing.T) {
	list := []int16{4, 2, 1, 3, 0, 9, 6, 8, 7, 5, 10}
	res := MergeInt16(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int16(i) {
			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
		}
	}
}

func TestMergeInt32(t *testing.T) {
	list := []int32{4, 2, 1, 3, 0, 9, 6, 8, 7, 5, 10}
	res := MergeInt32(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int32(i) {
			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
		}
	}
}

func TestMergeInt64(t *testing.T) {
	list := []int64{4, 2, 1, 3, 0, 9, 6, 8, 7, 5, 10}
	res := MergeInt64(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int64(i) {
			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
		}
	}
}

func TestMergeFloat32(t *testing.T) {
	list := []float32{4, 2, 1, 3, 0, 9, 6, 8, 7, 5, 10}
	res := MergeFloat32(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != float32(i) {
			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
		}
	}
}

func TestMergeFloat64(t *testing.T) {
	list := []float64{4, 2, 1, 3, 0, 9, 6, 8, 7, 5, 10}
	res := MergeFloat64(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != float64(i) {
			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
		}
	}
}

func TestMergeByte(t *testing.T) {
	list := []byte{4, 2, 1, 3, 0, 9, 6, 8, 7, 5, 10}
	res := MergeByte(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != byte(i) {
			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
		}
	}
}

func TestMergeRune(t *testing.T) {
	list := []rune{4, 2, 1, 3, 0, 9, 6, 8, 7, 5, 10}
	res := MergeRune(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != rune(i) {
			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
		}
	}
}

func TestMergeUintptr(t *testing.T) {
	list := []uintptr{4, 2, 1, 3, 0, 9, 6, 8, 7, 5, 10}
	res := MergeUintptr(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uintptr(i) {
			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
		}
	}
}