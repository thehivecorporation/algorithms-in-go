package sort

import (
	"testing"
	"math/rand"
)

func TestShellUint8(t *testing.T) {
	list := []uint8{4, 2, 1, 3, 0,9,6,8,7,5}
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

func TestShellInt(t *testing.T) {
	list := []int{4, 2, 1, 3, 0,9,6,8,7,5,11,10}
	res := ShellInt(list, 0)

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

// func TestShellUint16(t *testing.T) {
// 	list := []uint16{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := ShellUint16(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != uint16(i) {
// 			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestShellUint32(t *testing.T) {
// 	list := []uint32{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := ShellUint32(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != uint32(i) {
// 			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestShellUint64(t *testing.T) {
// 	list := []uint64{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := ShellUint64(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != uint64(i) {
// 			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestShellInt8(t *testing.T) {
// 	list := []int8{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := ShellInt8(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != int8(i) {
// 			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestShellInt16(t *testing.T) {
// 	list := []int16{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := ShellInt16(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != int16(i) {
// 			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestShellInt32(t *testing.T) {
// 	list := []int32{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := ShellInt32(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != int32(i) {
// 			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestShellInt64(t *testing.T) {
// 	list := []int64{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := ShellInt64(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != int64(i) {
// 			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestShellFloat32(t *testing.T) {
// 	list := []float32{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := ShellFloat32(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != float32(i) {
// 			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestShellFloat64(t *testing.T) {
// 	list := []float64{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := ShellFloat64(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != float64(i) {
// 			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestShellByte(t *testing.T) {
// 	list := []byte{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := ShellByte(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != byte(i) {
// 			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// func TestShellRune(t *testing.T) {
// 	list := []rune{4, 2, 1, 7, 3, 0, 5, 6}
// 	res := ShellRune(list)

// 	if len(res) != len(list) {
// 		t.Fatalf("Returned list has a different number of items %d!=%d",
// 			len(list), len(res))
// 	}

// 	for i := 0; i < len(list); i++ {
// 		if res[i] != rune(i) {
// 			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
// 		}
// 	}
// }

// /* -------------------------------------------------------------------------- */

 func BenchmarkShellInt(b *testing.B) {
 	rand.Seed(int64(53))
 	list := make([]int, b.N)
 	for i := 0; i < b.N; i++ {
 		list[i] = rand.Int()
 	}

 	ShellInt(list, 4)
 }

// func BenchmarkShellUint(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]uint, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = uint(rand.Uint32())
// 	}

// 	ShellUint(list)
// }

// func BenchmarkShellUintptr(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]uintptr, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = uintptr(rand.Uint32())
// 	}

// 	ShellUintptr(list)
// }

// func BenchmarkShellRune(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]rune, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = rune(rand.Int31())
// 	}

// 	ShellRune(list)
// }

// func BenchmarkShellByte(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]byte, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = byte(rand.Uint32())
// 	}

// 	ShellByte(list)
// }

// func BenchmarkShellFloat64(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]float64, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = float64(rand.Float64())
// 	}

// 	ShellFloat64(list)
// }

// func BenchmarkShellFloat32(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]float32, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = float32(rand.Float32())
// 	}

// 	ShellFloat32(list)
// }

// func BenchmarkShellInt32(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]int32, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = int32(rand.Int31())
// 	}

// 	ShellInt32(list)
// }

// func BenchmarkShellInt16(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]int16, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = int16(rand.Int31())
// 	}

// 	ShellInt16(list)
// }

// func BenchmarkShellInt8(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]int8, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = int8(rand.Int31())
// 	}

// 	ShellInt8(list)
// }

// func BenchmarkShellUint64(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]uint64, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = uint64(rand.Uint32())
// 	}

// 	ShellUint64(list)
// }

// func BenchmarkShellUint32(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]uint32, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = rand.Uint32()
// 	}

// 	ShellUint32(list)
// }

// func BenchmarkShellUint16(b *testing.B) {
// 	rand.Seed(int64(53))
// 	list := make([]uint16, b.N)
// 	for i := 0; i < b.N; i++ {
// 		list[i] = uint16(rand.Uint32())
// 	}

// 	ShellUint16(list)
// }

 func BenchmarkShellUint8(b *testing.B) {
 	rand.Seed(int64(53))
 	list := make([]uint8, b.N)
 	for i := 0; i < b.N; i++ {
 		list[i] = uint8(rand.Uint32())
 	}

 	ShellUint8(list, 4)
 }
