package sort

import (
	"math/rand"
	"testing"
)

func TestSelectionUint8(t *testing.T) {
	list := []uint8{4, 2, 1, 7, 3, 0, 5, 6}
	res := SelectionUint8(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint8(i) {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}

func TestSelectionUint16(t *testing.T) {
	list := []uint16{4, 2, 1, 7, 3, 0, 5, 6}
	res := SelectionUint16(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint16(i) {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}

func TestSelectionUint32(t *testing.T) {
	list := []uint32{4, 2, 1, 7, 3, 0, 5, 6}
	res := SelectionUint32(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint32(i) {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}

func TestSelectionUint64(t *testing.T) {
	list := []uint64{4, 2, 1, 7, 3, 0, 5, 6}
	res := SelectionUint64(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint64(i) {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}

func TestSelectionInt8(t *testing.T) {
	list := []int8{4, 2, 1, 7, 3, 0, 5, 6}
	res := SelectionInt8(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int8(i) {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}

func TestSelectionInt16(t *testing.T) {
	list := []int16{4, 2, 1, 7, 3, 0, 5, 6}
	res := SelectionInt16(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int16(i) {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}

func TestSelectionInt32(t *testing.T) {
	list := []int32{4, 2, 1, 7, 3, 0, 5, 6}
	res := SelectionInt32(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int32(i) {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}

func TestSelectionInt64(t *testing.T) {
	list := []int64{4, 2, 1, 7, 3, 0, 5, 6}
	res := SelectionInt64(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int64(i) {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}

func TestSelectionFloat32(t *testing.T) {
	list := []float32{4, 2, 1, 7, 3, 0, 5, 6}
	res := SelectionFloat32(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != float32(i) {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}

func TestSelectionFloat64(t *testing.T) {
	list := []float64{4, 2, 1, 7, 3, 0, 5, 6}
	res := SelectionFloat64(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != float64(i) {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}

func TestSelectionByte(t *testing.T) {
	list := []byte{4, 2, 1, 7, 3, 0, 5, 6}
	res := SelectionByte(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != byte(i) {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}

func TestSelectionRune(t *testing.T) {
	list := []rune{4, 2, 1, 7, 3, 0, 5, 6}
	res := SelectionRune(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != rune(i) {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}

/* -------------------------------------------------------------------------- */

func BenchmarkSelectionInt(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = rand.Int()
	}

	InsertionInt(list)
}

func BenchmarkSelectionUint(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uint, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = uint(rand.Uint32())
	}

	InsertionUint(list)
}

func BenchmarkSelectionUintptr(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uintptr, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = uintptr(rand.Uint32())
	}

	InsertionUintptr(list)
}

func BenchmarkSelectionRune(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]rune, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = rune(rand.Int31())
	}

	InsertionRune(list)
}

func BenchmarkSelectionByte(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]byte, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = byte(rand.Uint32())
	}

	InsertionByte(list)
}

func BenchmarkSelectionFloat64(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]float64, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = float64(rand.Float64())
	}

	InsertionFloat64(list)
}

func BenchmarkSelectionFloat32(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]float32, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = float32(rand.Float32())
	}

	InsertionFloat32(list)
}

func BenchmarkSelectionInt32(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]int32, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = int32(rand.Int31())
	}

	InsertionInt32(list)
}

func BenchmarkSelectionInt16(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]int16, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = int16(rand.Int31())
	}

	InsertionInt16(list)
}

func BenchmarkSelectionInt8(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]int8, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = int8(rand.Int31())
	}

	InsertionInt8(list)
}

func BenchmarkSelectionUint64(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uint64, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = uint64(rand.Uint32())
	}

	InsertionUint64(list)
}

func BenchmarkSelectionUint32(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uint32, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = rand.Uint32()
	}

	InsertionUint32(list)
}

func BenchmarkSelectionUint16(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uint16, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = uint16(rand.Uint32())
	}

	InsertionUint16(list)
}

func BenchmarkSelectionUint8(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uint8, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = uint8(rand.Uint32())
	}

	InsertionUint8(list)
}
