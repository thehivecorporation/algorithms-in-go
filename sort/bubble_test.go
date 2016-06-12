package sort

import (
	"math/rand"
	"testing"
)

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

//func TestBubbleInt8(t *testing.T) {
//	list := []int8{4, 2, 1, 7, 3, 0, 5, 6}
//	res := BubbleInt8(list)
//
//	if len(res) != len(list) {
//		t.Fatalf("Returned list has a different number of items %d!=%d",
//			len(list), len(res))
//	}
//
//	for i := 0; i < len(list); i++ {
//		if res[i] != int8(i) {
//			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
//		}
//	}
//}
//
//func TestBubbleInt16(t *testing.T) {
//	list := []int16{4, 2, 1, 7, 3, 0, 5, 6}
//	res := BubbleInt16(list)
//
//	if len(res) != len(list) {
//		t.Fatalf("Returned list has a different number of items %d!=%d",
//			len(list), len(res))
//	}
//
//	for i := 0; i < len(list); i++ {
//		if res[i] != int16(i) {
//			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
//		}
//	}
//}
//
//func TestBubbleInt32(t *testing.T) {
//	list := []int32{4, 2, 1, 7, 3, 0, 5, 6}
//	res := BubbleInt32(list)
//
//	if len(res) != len(list) {
//		t.Fatalf("Returned list has a different number of items %d!=%d",
//			len(list), len(res))
//	}
//
//	for i := 0; i < len(list); i++ {
//		if res[i] != int32(i) {
//			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
//		}
//	}
//}
//
//func TestBubbleInt64(t *testing.T) {
//	list := []int64{4, 2, 1, 7, 3, 0, 5, 6}
//	res := BubbleInt64(list)
//
//	if len(res) != len(list) {
//		t.Fatalf("Returned list has a different number of items %d!=%d",
//			len(list), len(res))
//	}
//
//	for i := 0; i < len(list); i++ {
//		if res[i] != int64(i) {
//			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
//		}
//	}
//}
//
//func TestBubbleFloat32(t *testing.T) {
//	list := []float32{4, 2, 1, 7, 3, 0, 5, 6}
//	res := BubbleFloat32(list)
//
//	if len(res) != len(list) {
//		t.Fatalf("Returned list has a different number of items %d!=%d",
//			len(list), len(res))
//	}
//
//	for i := 0; i < len(list); i++ {
//		if res[i] != float32(i) {
//			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
//		}
//	}
//}
//
//func TestBubbleFloat64(t *testing.T) {
//	list := []float64{4, 2, 1, 7, 3, 0, 5, 6}
//	res := BubbleFloat64(list)
//
//	if len(res) != len(list) {
//		t.Fatalf("Returned list has a different number of items %d!=%d",
//			len(list), len(res))
//	}
//
//	for i := 0; i < len(list); i++ {
//		if res[i] != float64(i) {
//			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
//		}
//	}
//}
//
//func TestBubbleByte(t *testing.T) {
//	list := []byte{4, 2, 1, 7, 3, 0, 5, 6}
//	res := BubbleByte(list)
//
//	if len(res) != len(list) {
//		t.Fatalf("Returned list has a different number of items %d!=%d",
//			len(list), len(res))
//	}
//
//	for i := 0; i < len(list); i++ {
//		if res[i] != byte(i) {
//			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
//		}
//	}
//}
//
//func TestBubbleRune(t *testing.T) {
//	list := []rune{4, 2, 1, 7, 3, 0, 5, 6}
//	res := BubbleRune(list)
//
//	if len(res) != len(list) {
//		t.Fatalf("Returned list has a different number of items %d!=%d",
//			len(list), len(res))
//	}
//
//	for i := 0; i < len(list); i++ {
//		if res[i] != rune(i) {
//			t.Fatalf("Bubble algorithm isn't correct: %v\n", res)
//		}
//	}
//}
//
///* -------------------------------------------------------------------------- */
//
//func BenchmarkBubbleInt(b *testing.B) {
//	rand.Seed(int64(53))
//	list := make([]int, b.N)
//	for i := 0; i < b.N; i++ {
//		list[i] = rand.Int()
//	}
//
//	BubbleInt(list)
//}
//
//func BenchmarkBubbleUint(b *testing.B) {
//	rand.Seed(int64(53))
//	list := make([]uint, b.N)
//	for i := 0; i < b.N; i++ {
//		list[i] = uint(rand.Uint32())
//	}
//
//	BubbleUint(list)
//}
//
//func BenchmarkBubbleUintptr(b *testing.B) {
//	rand.Seed(int64(53))
//	list := make([]uintptr, b.N)
//	for i := 0; i < b.N; i++ {
//		list[i] = uintptr(rand.Uint32())
//	}
//
//	BubbleUintptr(list)
//}
//
//func BenchmarkBubbleRune(b *testing.B) {
//	rand.Seed(int64(53))
//	list := make([]rune, b.N)
//	for i := 0; i < b.N; i++ {
//		list[i] = rune(rand.Int31())
//	}
//
//	BubbleRune(list)
//}
//
//func BenchmarkBubbleByte(b *testing.B) {
//	rand.Seed(int64(53))
//	list := make([]byte, b.N)
//	for i := 0; i < b.N; i++ {
//		list[i] = byte(rand.Uint32())
//	}
//
//	BubbleByte(list)
//}
//
//func BenchmarkBubbleFloat64(b *testing.B) {
//	rand.Seed(int64(53))
//	list := make([]float64, b.N)
//	for i := 0; i < b.N; i++ {
//		list[i] = float64(rand.Float64())
//	}
//
//	BubbleFloat64(list)
//}
//
//func BenchmarkBubbleFloat32(b *testing.B) {
//	rand.Seed(int64(53))
//	list := make([]float32, b.N)
//	for i := 0; i < b.N; i++ {
//		list[i] = float32(rand.Float32())
//	}
//
//	BubbleFloat32(list)
//}
//
//func BenchmarkBubbleInt32(b *testing.B) {
//	rand.Seed(int64(53))
//	list := make([]int32, b.N)
//	for i := 0; i < b.N; i++ {
//		list[i] = int32(rand.Int31())
//	}
//
//	BubbleInt32(list)
//}
//
//func BenchmarkBubbleInt16(b *testing.B) {
//	rand.Seed(int64(53))
//	list := make([]int16, b.N)
//	for i := 0; i < b.N; i++ {
//		list[i] = int16(rand.Int31())
//	}
//
//	BubbleInt16(list)
//}
//
//func BenchmarkBubbleInt8(b *testing.B) {
//	rand.Seed(int64(53))
//	list := make([]int8, b.N)
//	for i := 0; i < b.N; i++ {
//		list[i] = int8(rand.Int31())
//	}
//
//	BubbleInt8(list)
//}
//
//func BenchmarkBubbleUint64(b *testing.B) {
//	rand.Seed(int64(53))
//	list := make([]uint64, b.N)
//	for i := 0; i < b.N; i++ {
//		list[i] = uint64(rand.Uint32())
//	}
//
//	BubbleUint64(list)
//}
//
//func BenchmarkBubbleUint32(b *testing.B) {
//	rand.Seed(int64(53))
//	list := make([]uint32, b.N)
//	for i := 0; i < b.N; i++ {
//		list[i] = rand.Uint32()
//	}
//
//	BubbleUint32(list)
//}
//
//func BenchmarkBubbleUint16(b *testing.B) {
//	rand.Seed(int64(53))
//	list := make([]uint16, b.N)
//	for i := 0; i < b.N; i++ {
//		list[i] = uint16(rand.Uint32())
//	}
//
//	BubbleUint16(list)
//}

func BenchmarkBubbleUint8(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]uint8, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = uint8(rand.Uint32())
	}

	BubbleUint8(list)
}
