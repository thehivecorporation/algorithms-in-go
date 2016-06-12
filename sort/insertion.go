package sort

func InsertionUint8(u []uint8) []uint8 {

	newList := make([]uint8, len(u))
	copy(newList, u)
	for i := 1; i < len(newList); i++ {
		val := newList[i]

		for k := i - 1; k >= 0; k-- {
			prev := newList[k]

			if val < prev {
				newList[k+1] = prev
				newList[k] = val
			}
		}
	}

	return newList
}

func InsertionUint16(u []uint16) []uint16 {

	newList := make([]uint16, len(u))
	copy(newList, u)
	for i := 1; i < len(newList); i++ {
		val := newList[i]

		for k := i - 1; k >= 0; k-- {
			prev := newList[k]

			if val < prev {
				newList[k+1] = prev
				newList[k] = val
			}
		}
	}

	return newList
}

func InsertionUint32(u []uint32) []uint32 {

	newList := make([]uint32, len(u))
	copy(newList, u)
	for i := 1; i < len(newList); i++ {
		val := newList[i]

		for k := i - 1; k >= 0; k-- {
			prev := newList[k]

			if val < prev {
				newList[k+1] = prev
				newList[k] = val
			}
		}
	}

	return newList
}

func InsertionUint64(u []uint64) []uint64 {

	newList := make([]uint64, len(u))
	copy(newList, u)
	for i := 1; i < len(newList); i++ {
		val := newList[i]

		for k := i - 1; k >= 0; k-- {
			prev := newList[k]

			if val < prev {
				newList[k+1] = prev
				newList[k] = val
			}
		}
	}

	return newList
}

func InsertionInt8(u []int8) []int8 {

	newList := make([]int8, len(u))
	copy(newList, u)
	for i := 1; i < len(newList); i++ {
		val := newList[i]

		for k := i - 1; k >= 0; k-- {
			prev := newList[k]

			if val < prev {
				newList[k+1] = prev
				newList[k] = val
			}
		}
	}

	return newList
}

func InsertionInt16(u []int16) []int16 {

	newList := make([]int16, len(u))
	copy(newList, u)
	for i := 1; i < len(newList); i++ {
		val := newList[i]

		for k := i - 1; k >= 0; k-- {
			prev := newList[k]

			if val < prev {
				newList[k+1] = prev
				newList[k] = val
			}
		}
	}

	return newList
}

func InsertionInt32(u []int32) []int32 {

	newList := make([]int32, len(u))
	copy(newList, u)
	for i := 1; i < len(newList); i++ {
		val := newList[i]

		for k := i - 1; k >= 0; k-- {
			prev := newList[k]

			if val < prev {
				newList[k+1] = prev
				newList[k] = val
			}
		}
	}

	return newList
}

func InsertionInt64(u []int64) []int64 {

	newList := make([]int64, len(u))
	copy(newList, u)
	for i := 1; i < len(newList); i++ {
		val := newList[i]

		for k := i - 1; k >= 0; k-- {
			prev := newList[k]

			if val < prev {
				newList[k+1] = prev
				newList[k] = val
			}
		}
	}

	return newList
}

func InsertionFloat32(u []float32) []float32 {

	newList := make([]float32, len(u))
	copy(newList, u)
	for i := 1; i < len(newList); i++ {
		val := newList[i]

		for k := i - 1; k >= 0; k-- {
			prev := newList[k]

			if val < prev {
				newList[k+1] = prev
				newList[k] = val
			}
		}
	}

	return newList
}

func InsertionFloat64(u []float64) []float64 {

	newList := make([]float64, len(u))
	copy(newList, u)
	for i := 1; i < len(newList); i++ {
		val := newList[i]

		for k := i - 1; k >= 0; k-- {
			prev := newList[k]

			if val < prev {
				newList[k+1] = prev
				newList[k] = val
			}
		}
	}

	return newList
}

func InsertionByte(u []byte) []byte {

	newList := make([]byte, len(u))
	copy(newList, u)
	for i := 1; i < len(newList); i++ {
		val := newList[i]

		for k := i - 1; k >= 0; k-- {
			prev := newList[k]

			if val < prev {
				newList[k+1] = prev
				newList[k] = val
			}
		}
	}

	return newList
}

func InsertionRune(u []rune) []rune {

	newList := make([]rune, len(u))
	copy(newList, u)
	for i := 1; i < len(newList); i++ {
		val := newList[i]

		for k := i - 1; k >= 0; k-- {
			prev := newList[k]

			if val < prev {
				newList[k+1] = prev
				newList[k] = val
			}
		}
	}

	return newList
}

func InsertionUint(u []uint) []uint {

	newList := make([]uint, len(u))
	copy(newList, u)
	for i := 1; i < len(newList); i++ {
		val := newList[i]

		for k := i - 1; k >= 0; k-- {
			prev := newList[k]

			if val < prev {
				newList[k+1] = prev
				newList[k] = val
			}
		}
	}

	return newList
}

func InsertionInt(u []int) []int {
	newList := make([]int, len(u))
	copy(newList, u)

	for i := 1; i < len(newList); i++ {
		val := newList[i]

		for k := i - 1; k >= 0; k-- {
			prev := newList[k]

			if val < prev {
				newList[k+1] = prev
				newList[k] = val
			}
		}
	}

	return newList
}

func InsertionUintptr(u []uintptr) []uintptr {

	newList := make([]uintptr, len(u))
	copy(newList, u)
	for i := 1; i < len(newList); i++ {
		val := newList[i]

		for k := i - 1; k >= 0; k-- {
			prev := newList[k]

			if val < prev {
				newList[k+1] = prev
				newList[k] = val
			}
		}
	}

	return newList
}
