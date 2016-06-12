package sort

import (
	"math/rand"
	"testing"
)

func TestInsertionInt(t *testing.T) {
	list := []int{4, 2, 1, 7, 5, 0, 3, 9, 6, 8}
	res := InsertionInt(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != i {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}

func TestInsertionUint(t *testing.T) {
	list := []uint{4, 2, 1, 7, 5, 0, 3, 9, 6, 8}
	res := InsertionUint(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint(i) {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}

func TestInsertionUintptr(t *testing.T) {
	list := []uintptr{4, 2, 1, 7, 5, 0, 3, 9, 6, 8}
	res := InsertionUintptr(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uintptr(i) {
			t.Fatalf("Insertion algorithm isn't correct: %v\n", res)
		}
	}
}

func TestInsertionRune(t *testing.T) {
	list := []rune{4, 2, 1, 7, 5, 0, 3, 9, 6, 8}
	res := InsertionRune(list)

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

func TestInsertionByte(t *testing.T) {
	list := []byte{4, 2, 1, 7, 5, 0, 3, 9, 6, 8}
	res := InsertionByte(list)

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

func TestInsertionFloat64(t *testing.T) {
	list := []float64{4, 2, 1, 7, 5, 0, 3, 9, 6, 8}
	res := InsertionFloat64(list)

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

func TestInsertionFloat32(t *testing.T) {
	list := []float32{4, 2, 1, 7, 5, 0, 3, 9, 6, 8}
	res := InsertionFloat32(list)

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

func TestInsertionInt64(t *testing.T) {
	list := []int64{4, 2, 1, 7, 5, 0, 3, 9, 6, 8}
	res := InsertionInt64(list)

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

func TestInsertionInt16(t *testing.T) {
	list := []int16{4, 2, 1, 7, 5, 0, 3, 9, 6, 8}
	res := InsertionInt16(list)

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

func TestInsertionInt8(t *testing.T) {
	list := []int8{4, 2, 1, 7, 5, 0, 3, 9, 6, 8}
	res := InsertionInt8(list)

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

func TestInsertionUint64(t *testing.T) {
	list := []uint64{4, 2, 1, 7, 5, 0, 3, 9, 6, 8}
	res := InsertionUint64(list)

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

func TestInsertionUint32(t *testing.T) {
	list := []uint32{4, 2, 1, 7, 5, 0, 3, 9, 6, 8}
	res := InsertionUint32(list)

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

func TestInsertionUint16(t *testing.T) {
	list := []uint16{4, 2, 1, 7, 5, 0, 3, 9, 6, 8}
	res := InsertionUint16(list)

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

func TestInsertionUint8(t *testing.T) {
	list := []uint8{4, 2, 1, 7, 5, 0, 3, 9, 6, 8}
	res := InsertionUint8(list)

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

/* -------------------------------------------------------------------------- */

func BenchmarkInsertionInt(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = rand.Int()
	}

	InsertionInt(list)
}

func BenchmarkInsertionUint(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uint, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = uint(rand.Uint32())
	}

	InsertionUint(list)
}

func BenchmarkInsertionUintptr(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uintptr, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = uintptr(rand.Uint32())
	}

	InsertionUintptr(list)
}

func BenchmarkInsertionRune(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]rune, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = rune(rand.Int31())
	}

	InsertionRune(list)
}

func BenchmarkInsertionByte(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]byte, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = byte(rand.Uint32())
	}

	InsertionByte(list)
}

func BenchmarkInsertionFloat64(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]float64, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = float64(rand.Float64())
	}

	InsertionFloat64(list)
}

func BenchmarkInsertionFloat32(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]float32, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = float32(rand.Float32())
	}

	InsertionFloat32(list)
}

func BenchmarkInsertionInt32(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]int32, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = int32(rand.Int31())
	}

	InsertionInt32(list)
}

func BenchmarkInsertionInt16(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]int16, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = int16(rand.Int31())
	}

	InsertionInt16(list)
}

func BenchmarkInsertionInt8(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]int8, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = int8(rand.Int31())
	}

	InsertionInt8(list)
}

func BenchmarkInsertionUint64(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uint64, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = uint64(rand.Uint32())
	}

	InsertionUint64(list)
}

func BenchmarkInsertionUint32(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uint32, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = rand.Uint32()
	}

	InsertionUint32(list)
}

func BenchmarkInsertionUint16(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uint16, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = uint16(rand.Uint32())
	}

	InsertionUint16(list)
}

func BenchmarkInsertionUint8(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uint8, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = uint8(rand.Uint32())
	}

	InsertionUint8(list)
}

func BenchmarkHello(b *testing.B) {
	hello(b.N)
}
