// DO NOT EDIT THIS FILE MANUALLY
// Generated by go generate
package sorting

import "testing"

func TestHeapUint(t *testing.T) {
	list := []uint{4, 2, 1, 3, 0}
	res := heapUint(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint(i) {
			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
		}
	}
}

func TestHeapUint8(t *testing.T) {
	list := []uint8{4, 2, 1, 3, 0}
	res := heapUint8(list)

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

func TestHeapUint16(t *testing.T) {
	list := []uint16{4, 2, 1, 3, 0}
	res := heapUint16(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint16(i) {
			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
		}
	}
}

func TestHeapUint32(t *testing.T) {
	list := []uint32{4, 2, 1, 3, 0}
	res := heapUint32(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint32(i) {
			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
		}
	}
}

func TestHeapUint64(t *testing.T) {
	list := []uint64{4, 2, 1, 3, 0}
	res := heapUint64(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint64(i) {
			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
		}
	}
}

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

func TestHeapInt8(t *testing.T) {
	list := []int8{4, 2, 1, 3, 0}
	res := heapInt8(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int8(i) {
			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
		}
	}
}

func TestHeapInt16(t *testing.T) {
	list := []int16{4, 2, 1, 3, 0}
	res := heapInt16(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int16(i) {
			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
		}
	}
}

func TestHeapInt32(t *testing.T) {
	list := []int32{4, 2, 1, 3, 0}
	res := heapInt32(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int32(i) {
			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
		}
	}
}

func TestHeapInt64(t *testing.T) {
	list := []int64{4, 2, 1, 3, 0}
	res := heapInt64(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int64(i) {
			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
		}
	}
}

func TestHeapFloat32(t *testing.T) {
	list := []float32{4, 2, 1, 3, 0}
	res := heapFloat32(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != float32(i) {
			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
		}
	}
}

func TestHeapFloat64(t *testing.T) {
	list := []float64{4, 2, 1, 3, 0}
	res := heapFloat64(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != float64(i) {
			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
		}
	}
}

func TestHeapByte(t *testing.T) {
	list := []byte{4, 2, 1, 3, 0}
	res := heapByte(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != byte(i) {
			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
		}
	}
}

func TestHeapRune(t *testing.T) {
	list := []rune{4, 2, 1, 3, 0}
	res := heapRune(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != rune(i) {
			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
		}
	}
}

func TestHeapUintptr(t *testing.T) {
	list := []uintptr{4, 2, 1, 3, 0}
	res := heapUintptr(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uintptr(i) {
			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
		}
	}
}
