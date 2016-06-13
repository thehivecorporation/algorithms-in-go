package sort

import (
	"math/rand"
	"testing"
)

//func TestHeapUint8(t *testing.T) {
//	list := []uint8{4, 2, 1, 3, 0, 9, 6, 8, 7, 5}
//	res := HeapUint8(list, 0)
//
//	if len(res) != len(list) {
//		t.Fatalf("Returned list has a different number of items %d!=%d",
//			len(list), len(res))
//	}
//
//	for i := 0; i < len(list); i++ {
//		if res[i] != uint8(i) {
//			t.Fatalf("Heap algorithm isn't correct: %v\n", res)
//		}
//	}
//}

func TestHeapInt(t *testing.T) {
	list := []int{4, 2, 1, 3, 0, 9, 6, 8, 7, 5, 10}
	res := HeapInt(list)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != int(i) {
			t.Fatalf("Heap algorithm isn't correct: %v\n", res)
		}
	}
}

// func TestHeapUint16(t *testing.T) {
// 	list := []uint16{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := HeapUint16(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != uint16(i) {
// 			t.Fatalf("Heap algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestHeapUint32(t *testing.T) {
// 	list := []uint32{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := HeapUint32(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != uint32(i) {
// 			t.Fatalf("Heap algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestHeapUint64(t *testing.T) {
// 	list := []uint64{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := HeapUint64(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != uint64(i) {
// 			t.Fatalf("Heap algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestHeapInt8(t *testing.T) {
// 	list := []int8{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := HeapInt8(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != int8(i) {
// 			t.Fatalf("Heap algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestHeapInt16(t *testing.T) {
// 	list := []int16{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := HeapInt16(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != int16(i) {
// 			t.Fatalf("Heap algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestHeapInt32(t *testing.T) {
// 	list := []int32{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := HeapInt32(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != int32(i) {
// 			t.Fatalf("Heap algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestHeapInt64(t *testing.T) {
// 	list := []int64{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := HeapInt64(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != int64(i) {
// 			t.Fatalf("Heap algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestHeapFloat32(t *testing.T) {
// 	list := []float32{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := HeapFloat32(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != float32(i) {
// 			t.Fatalf("Heap algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestHeapFloat64(t *testing.T) {
// 	list := []float64{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := HeapFloat64(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != float64(i) {
// 			t.Fatalf("Heap algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestHeapByte(t *testing.T) {
// 	list := []byte{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := HeapByte(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != byte(i) {
// 			t.Fatalf("Heap algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestHeapRune(t *testing.T) {
// 	list := []rune{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := HeapRune(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != rune(i) {
// 			t.Fatalf("Heap algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// /* -------------------------------------------------------------------------- */

func BenchmarkHeapInt(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = rand.Int()
	}

	HeapInt(list)
}


// func BenchmarkHeapUint(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]uint, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = uint(rand.Uint32())
// 	}

// 	HeapUint(list)
// }

// func BenchmarkHeapUintptr(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]uintptr, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = uintptr(rand.Uint32())
// 	}

// 	HeapUintptr(list)
// }

// func BenchmarkHeapRune(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]rune, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = rune(rand.Int31())
// 	}

// 	HeapRune(list)
// }

// func BenchmarkHeapByte(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]byte, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = byte(rand.Uint32())
// 	}

// 	HeapByte(list)
// }

// func BenchmarkHeapFloat64(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]float64, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = float64(rand.Float64())
// 	}

// 	HeapFloat64(list)
// }

// func BenchmarkHeapFloat32(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]float32, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = float32(rand.Float32())
// 	}

// 	HeapFloat32(list)
// }

// func BenchmarkHeapInt32(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]int32, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = int32(rand.Int31())
// 	}

// 	HeapInt32(list)
// }

// func BenchmarkHeapInt16(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]int16, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = int16(rand.Int31())
// 	}

// 	HeapInt16(list)
// }

// func BenchmarkHeapInt8(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]int8, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = int8(rand.Int31())
// 	}

// 	HeapInt8(list)
// }

// func BenchmarkHeapUint64(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]uint64, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = uint64(rand.Uint32())
// 	}

// 	HeapUint64(list)
// }

// func BenchmarkHeapUint32(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]uint32, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = rand.Uint32()
// 	}

// 	HeapUint32(list)
// }

// func BenchmarkHeapUint16(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]uint16, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = uint16(rand.Uint32())
// 	}

// 	HeapUint16(list)
// }

//func BenchmarkHeapUint8(b *testing.B) {
//	rand.Seed(int64(53))
//	list := make([]uint8, b.N)
//	for i := 0; i < b.N; i++ {
//		list[i] = uint8(rand.Uint32())
//	}
//
//	HeapUint8(list, 4)
//}
