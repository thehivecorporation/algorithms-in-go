// DO NOT EDIT THIS FILE MANUALLY
// Generated by go generate
package sorting

import "testing"

func TestQuickUint(t *testing.T) {
	list := []uint{4, 2, 1, 3, 0, 5}
	list2 := make([]uint, len(list))
	copy(list2, list)

	res := QuickUint(list)

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
		if res[i] != uint(i) {
			t.Fatalf("Quicksort algorithm isn't correct: %v\n", res)
		}
	}
}

func TestQuickUint8(t *testing.T) {
	list := []uint8{4, 2, 1, 3, 0, 5}
	list2 := make([]uint8, len(list))
	copy(list2, list)

	res := QuickUint8(list)

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
		if res[i] != uint8(i) {
			t.Fatalf("Quicksort algorithm isn't correct: %v\n", res)
		}
	}
}

func TestQuickUint16(t *testing.T) {
	list := []uint16{4, 2, 1, 3, 0, 5}
	list2 := make([]uint16, len(list))
	copy(list2, list)

	res := QuickUint16(list)

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
		if res[i] != uint16(i) {
			t.Fatalf("Quicksort algorithm isn't correct: %v\n", res)
		}
	}
}

func TestQuickUint32(t *testing.T) {
	list := []uint32{4, 2, 1, 3, 0, 5}
	list2 := make([]uint32, len(list))
	copy(list2, list)

	res := QuickUint32(list)

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
		if res[i] != uint32(i) {
			t.Fatalf("Quicksort algorithm isn't correct: %v\n", res)
		}
	}
}

func TestQuickUint64(t *testing.T) {
	list := []uint64{4, 2, 1, 3, 0, 5}
	list2 := make([]uint64, len(list))
	copy(list2, list)

	res := QuickUint64(list)

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
		if res[i] != uint64(i) {
			t.Fatalf("Quicksort algorithm isn't correct: %v\n", res)
		}
	}
}

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

func TestQuickInt8(t *testing.T) {
	list := []int8{4, 2, 1, 3, 0, 5}
	list2 := make([]int8, len(list))
	copy(list2, list)

	res := QuickInt8(list)

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
		if res[i] != int8(i) {
			t.Fatalf("Quicksort algorithm isn't correct: %v\n", res)
		}
	}
}

func TestQuickInt16(t *testing.T) {
	list := []int16{4, 2, 1, 3, 0, 5}
	list2 := make([]int16, len(list))
	copy(list2, list)

	res := QuickInt16(list)

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
		if res[i] != int16(i) {
			t.Fatalf("Quicksort algorithm isn't correct: %v\n", res)
		}
	}
}

func TestQuickInt32(t *testing.T) {
	list := []int32{4, 2, 1, 3, 0, 5}
	list2 := make([]int32, len(list))
	copy(list2, list)

	res := QuickInt32(list)

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
		if res[i] != int32(i) {
			t.Fatalf("Quicksort algorithm isn't correct: %v\n", res)
		}
	}
}

func TestQuickInt64(t *testing.T) {
	list := []int64{4, 2, 1, 3, 0, 5}
	list2 := make([]int64, len(list))
	copy(list2, list)

	res := QuickInt64(list)

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
		if res[i] != int64(i) {
			t.Fatalf("Quicksort algorithm isn't correct: %v\n", res)
		}
	}
}

func TestQuickFloat32(t *testing.T) {
	list := []float32{4, 2, 1, 3, 0, 5}
	list2 := make([]float32, len(list))
	copy(list2, list)

	res := QuickFloat32(list)

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
		if res[i] != float32(i) {
			t.Fatalf("Quicksort algorithm isn't correct: %v\n", res)
		}
	}
}

func TestQuickFloat64(t *testing.T) {
	list := []float64{4, 2, 1, 3, 0, 5}
	list2 := make([]float64, len(list))
	copy(list2, list)

	res := QuickFloat64(list)

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
		if res[i] != float64(i) {
			t.Fatalf("Quicksort algorithm isn't correct: %v\n", res)
		}
	}
}

func TestQuickByte(t *testing.T) {
	list := []byte{4, 2, 1, 3, 0, 5}
	list2 := make([]byte, len(list))
	copy(list2, list)

	res := QuickByte(list)

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
		if res[i] != byte(i) {
			t.Fatalf("Quicksort algorithm isn't correct: %v\n", res)
		}
	}
}

func TestQuickRune(t *testing.T) {
	list := []rune{4, 2, 1, 3, 0, 5}
	list2 := make([]rune, len(list))
	copy(list2, list)

	res := QuickRune(list)

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
		if res[i] != rune(i) {
			t.Fatalf("Quicksort algorithm isn't correct: %v\n", res)
		}
	}
}

func TestQuickUintptr(t *testing.T) {
	list := []uintptr{4, 2, 1, 3, 0, 5}
	list2 := make([]uintptr, len(list))
	copy(list2, list)

	res := QuickUintptr(list)

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
		if res[i] != uintptr(i) {
			t.Fatalf("Quicksort algorithm isn't correct: %v\n", res)
		}
	}
}
