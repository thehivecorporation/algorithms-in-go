// DO NOT EDIT THIS FILE MANUALLY
// Generated by go generate
package sorting

import (
	"math/rand"
	"testing"
)

func TestBubbleUint(t *testing.T) {
	list := []uint{4, 2, 1, 7, 3, 0, 5, 6}
	res := BubbleUint(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint(i) {
			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
		}
	}
}

func BenchmarkBubbleUint(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uint, 1000)
	for i := 0; i < 1000; i++ {
		list[i] = uint(rand.Int63())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BubbleUint(list)
	}
}

func TestBubbleUint8(t *testing.T) {
	list := []uint8{4, 2, 1, 7, 3, 0, 5, 6}
	res := BubbleUint8(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint8(i) {
			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
		}
	}
}

func BenchmarkBubbleUint8(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uint8, 1000)
	for i := 0; i < 1000; i++ {
		list[i] = uint8(rand.Int63())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BubbleUint8(list)
	}
}

func TestBubbleUint16(t *testing.T) {
	list := []uint16{4, 2, 1, 7, 3, 0, 5, 6}
	res := BubbleUint16(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint16(i) {
			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
		}
	}
}

func BenchmarkBubbleUint16(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uint16, 1000)
	for i := 0; i < 1000; i++ {
		list[i] = uint16(rand.Int63())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BubbleUint16(list)
	}
}

func TestBubbleUint32(t *testing.T) {
	list := []uint32{4, 2, 1, 7, 3, 0, 5, 6}
	res := BubbleUint32(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint32(i) {
			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
		}
	}
}

func BenchmarkBubbleUint32(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uint32, 1000)
	for i := 0; i < 1000; i++ {
		list[i] = uint32(rand.Int63())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BubbleUint32(list)
	}
}

func TestBubbleUint64(t *testing.T) {
	list := []uint64{4, 2, 1, 7, 3, 0, 5, 6}
	res := BubbleUint64(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uint64(i) {
			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
		}
	}
}

func BenchmarkBubbleUint64(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uint64, 1000)
	for i := 0; i < 1000; i++ {
		list[i] = uint64(rand.Int63())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BubbleUint64(list)
	}
}

func TestBubbleInt(t *testing.T) {
	list := []int{4, 2, 1, 7, 3, 0, 5, 6}
	res := BubbleInt(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int(i) {
			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
		}
	}
}

func BenchmarkBubbleInt(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		list[i] = int(rand.Int63())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BubbleInt(list)
	}
}

func TestBubbleInt8(t *testing.T) {
	list := []int8{4, 2, 1, 7, 3, 0, 5, 6}
	res := BubbleInt8(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int8(i) {
			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
		}
	}
}

func BenchmarkBubbleInt8(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]int8, 1000)
	for i := 0; i < 1000; i++ {
		list[i] = int8(rand.Int63())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BubbleInt8(list)
	}
}

func TestBubbleInt16(t *testing.T) {
	list := []int16{4, 2, 1, 7, 3, 0, 5, 6}
	res := BubbleInt16(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int16(i) {
			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
		}
	}
}

func BenchmarkBubbleInt16(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]int16, 1000)
	for i := 0; i < 1000; i++ {
		list[i] = int16(rand.Int63())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BubbleInt16(list)
	}
}

func TestBubbleInt32(t *testing.T) {
	list := []int32{4, 2, 1, 7, 3, 0, 5, 6}
	res := BubbleInt32(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int32(i) {
			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
		}
	}
}

func BenchmarkBubbleInt32(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]int32, 1000)
	for i := 0; i < 1000; i++ {
		list[i] = int32(rand.Int63())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BubbleInt32(list)
	}
}

func TestBubbleInt64(t *testing.T) {
	list := []int64{4, 2, 1, 7, 3, 0, 5, 6}
	res := BubbleInt64(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int64(i) {
			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
		}
	}
}

func BenchmarkBubbleInt64(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]int64, 1000)
	for i := 0; i < 1000; i++ {
		list[i] = int64(rand.Int63())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BubbleInt64(list)
	}
}

func TestBubbleFloat32(t *testing.T) {
	list := []float32{4, 2, 1, 7, 3, 0, 5, 6}
	res := BubbleFloat32(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != float32(i) {
			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
		}
	}
}

func BenchmarkBubbleFloat32(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]float32, 1000)
	for i := 0; i < 1000; i++ {
		list[i] = float32(rand.Int63())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BubbleFloat32(list)
	}
}

func TestBubbleFloat64(t *testing.T) {
	list := []float64{4, 2, 1, 7, 3, 0, 5, 6}
	res := BubbleFloat64(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != float64(i) {
			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
		}
	}
}

func BenchmarkBubbleFloat64(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]float64, 1000)
	for i := 0; i < 1000; i++ {
		list[i] = float64(rand.Int63())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BubbleFloat64(list)
	}
}

func TestBubbleByte(t *testing.T) {
	list := []byte{4, 2, 1, 7, 3, 0, 5, 6}
	res := BubbleByte(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != byte(i) {
			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
		}
	}
}

func BenchmarkBubbleByte(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]byte, 1000)
	for i := 0; i < 1000; i++ {
		list[i] = byte(rand.Int63())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BubbleByte(list)
	}
}

func TestBubbleRune(t *testing.T) {
	list := []rune{4, 2, 1, 7, 3, 0, 5, 6}
	res := BubbleRune(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != rune(i) {
			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
		}
	}
}

func BenchmarkBubbleRune(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]rune, 1000)
	for i := 0; i < 1000; i++ {
		list[i] = rune(rand.Int63())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BubbleRune(list)
	}
}

func TestBubbleUintptr(t *testing.T) {
	list := []uintptr{4, 2, 1, 7, 3, 0, 5, 6}
	res := BubbleUintptr(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != uintptr(i) {
			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
		}
	}
}

func BenchmarkBubbleUintptr(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uintptr, 1000)
	for i := 0; i < 1000; i++ {
		list[i] = uintptr(rand.Int63())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BubbleUintptr(list)
	}
}