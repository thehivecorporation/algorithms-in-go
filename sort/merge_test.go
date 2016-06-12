package sort

import (
	"math/rand"
	"testing"
)

//func TestMergeUint8(t *testing.T) {
//	list := []uint8{4, 2, 1, 3, 0, 9, 6, 8, 7, 5}
//	res := MergeUint8(list, 0)
//
//	if len(res) != len(list) {
//		t.Fatalf("Returned list has a different number of items %d!=%d",
//			len(list), len(res))
//	}
//
//	for i := 0; i < len(list); i++ {
//		if res[i] != uint8(i) {
//			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
//		}
//	}
//}

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

// func TestMergeUint16(t *testing.T) {
// 	list := []uint16{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := MergeUint16(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != uint16(i) {
// 			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestMergeUint32(t *testing.T) {
// 	list := []uint32{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := MergeUint32(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != uint32(i) {
// 			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestMergeUint64(t *testing.T) {
// 	list := []uint64{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := MergeUint64(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != uint64(i) {
// 			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestMergeInt8(t *testing.T) {
// 	list := []int8{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := MergeInt8(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != int8(i) {
// 			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestMergeInt16(t *testing.T) {
// 	list := []int16{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := MergeInt16(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != int16(i) {
// 			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestMergeInt32(t *testing.T) {
// 	list := []int32{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := MergeInt32(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != int32(i) {
// 			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestMergeInt64(t *testing.T) {
// 	list := []int64{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := MergeInt64(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != int64(i) {
// 			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestMergeFloat32(t *testing.T) {
// 	list := []float32{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := MergeFloat32(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != float32(i) {
// 			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestMergeFloat64(t *testing.T) {
// 	list := []float64{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := MergeFloat64(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != float64(i) {
// 			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestMergeByte(t *testing.T) {
// 	list := []byte{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := MergeByte(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != byte(i) {
// 			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestMergeRune(t *testing.T) {
// 	list := []rune{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := MergeRune(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != rune(i) {
// 			t.Fatalf("Merge algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// /* -------------------------------------------------------------------------- */

func BenchmarkMergeInt(b *testing.B) {
	rand.Seed(int64(53))
	list := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		list[i] = rand.Int()
	}

	MergeInt(list)
}


// func BenchmarkMergeUint(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]uint, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = uint(rand.Uint32())
// 	}

// 	MergeUint(list)
// }

// func BenchmarkMergeUintptr(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]uintptr, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = uintptr(rand.Uint32())
// 	}

// 	MergeUintptr(list)
// }

// func BenchmarkMergeRune(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]rune, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = rune(rand.Int31())
// 	}

// 	MergeRune(list)
// }

// func BenchmarkMergeByte(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]byte, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = byte(rand.Uint32())
// 	}

// 	MergeByte(list)
// }

// func BenchmarkMergeFloat64(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]float64, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = float64(rand.Float64())
// 	}

// 	MergeFloat64(list)
// }

// func BenchmarkMergeFloat32(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]float32, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = float32(rand.Float32())
// 	}

// 	MergeFloat32(list)
// }

// func BenchmarkMergeInt32(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]int32, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = int32(rand.Int31())
// 	}

// 	MergeInt32(list)
// }

// func BenchmarkMergeInt16(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]int16, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = int16(rand.Int31())
// 	}

// 	MergeInt16(list)
// }

// func BenchmarkMergeInt8(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]int8, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = int8(rand.Int31())
// 	}

// 	MergeInt8(list)
// }

// func BenchmarkMergeUint64(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]uint64, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = uint64(rand.Uint32())
// 	}

// 	MergeUint64(list)
// }

// func BenchmarkMergeUint32(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]uint32, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = rand.Uint32()
// 	}

// 	MergeUint32(list)
// }

// func BenchmarkMergeUint16(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]uint16, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = uint16(rand.Uint32())
// 	}

// 	MergeUint16(list)
// }

//func BenchmarkMergeUint8(b *testing.B) {
//	rand.Seed(int64(53))
//	list := make([]uint8, b.N)
//	for i := 0; i < b.N; i++ {
//		list[i] = uint8(rand.Uint32())
//	}
//
//	MergeUint8(list, 4)
//}
