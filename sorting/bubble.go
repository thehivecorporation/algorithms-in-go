// DO NOT EDIT THIS FILE MANUALLY
// Generated by go generate
package sorting

func BubbleUint(u []uint) (newList []uint) {
	newList = make([]uint, len(u))
	copy(newList, u)

	for {
		hasSwap := false
		for i := 0; i < len(newList)-1; i++ {
			if newList[i] > newList[i+1] {
				old := newList[i]
				newList[i] = newList[i+1]
				newList[i+1] = old
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}

	return newList
}

func BubbleUint8(u []uint8) (newList []uint8) {
	newList = make([]uint8, len(u))
	copy(newList, u)

	for {
		hasSwap := false
		for i := 0; i < len(newList)-1; i++ {
			if newList[i] > newList[i+1] {
				old := newList[i]
				newList[i] = newList[i+1]
				newList[i+1] = old
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}

	return newList
}

func BubbleUint16(u []uint16) (newList []uint16) {
	newList = make([]uint16, len(u))
	copy(newList, u)

	for {
		hasSwap := false
		for i := 0; i < len(newList)-1; i++ {
			if newList[i] > newList[i+1] {
				old := newList[i]
				newList[i] = newList[i+1]
				newList[i+1] = old
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}

	return newList
}

func BubbleUint32(u []uint32) (newList []uint32) {
	newList = make([]uint32, len(u))
	copy(newList, u)

	for {
		hasSwap := false
		for i := 0; i < len(newList)-1; i++ {
			if newList[i] > newList[i+1] {
				old := newList[i]
				newList[i] = newList[i+1]
				newList[i+1] = old
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}

	return newList
}

func BubbleUint64(u []uint64) (newList []uint64) {
	newList = make([]uint64, len(u))
	copy(newList, u)

	for {
		hasSwap := false
		for i := 0; i < len(newList)-1; i++ {
			if newList[i] > newList[i+1] {
				old := newList[i]
				newList[i] = newList[i+1]
				newList[i+1] = old
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}

	return newList
}

func BubbleInt(u []int) (newList []int) {
	newList = make([]int, len(u))
	copy(newList, u)

	for {
		hasSwap := false
		for i := 0; i < len(newList)-1; i++ {
			if newList[i] > newList[i+1] {
				old := newList[i]
				newList[i] = newList[i+1]
				newList[i+1] = old
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}

	return newList
}

func BubbleInt8(u []int8) (newList []int8) {
	newList = make([]int8, len(u))
	copy(newList, u)

	for {
		hasSwap := false
		for i := 0; i < len(newList)-1; i++ {
			if newList[i] > newList[i+1] {
				old := newList[i]
				newList[i] = newList[i+1]
				newList[i+1] = old
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}

	return newList
}

func BubbleInt16(u []int16) (newList []int16) {
	newList = make([]int16, len(u))
	copy(newList, u)

	for {
		hasSwap := false
		for i := 0; i < len(newList)-1; i++ {
			if newList[i] > newList[i+1] {
				old := newList[i]
				newList[i] = newList[i+1]
				newList[i+1] = old
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}

	return newList
}

func BubbleInt32(u []int32) (newList []int32) {
	newList = make([]int32, len(u))
	copy(newList, u)

	for {
		hasSwap := false
		for i := 0; i < len(newList)-1; i++ {
			if newList[i] > newList[i+1] {
				old := newList[i]
				newList[i] = newList[i+1]
				newList[i+1] = old
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}

	return newList
}

func BubbleInt64(u []int64) (newList []int64) {
	newList = make([]int64, len(u))
	copy(newList, u)

	for {
		hasSwap := false
		for i := 0; i < len(newList)-1; i++ {
			if newList[i] > newList[i+1] {
				old := newList[i]
				newList[i] = newList[i+1]
				newList[i+1] = old
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}

	return newList
}

func BubbleFloat32(u []float32) (newList []float32) {
	newList = make([]float32, len(u))
	copy(newList, u)

	for {
		hasSwap := false
		for i := 0; i < len(newList)-1; i++ {
			if newList[i] > newList[i+1] {
				old := newList[i]
				newList[i] = newList[i+1]
				newList[i+1] = old
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}

	return newList
}

func BubbleFloat64(u []float64) (newList []float64) {
	newList = make([]float64, len(u))
	copy(newList, u)

	for {
		hasSwap := false
		for i := 0; i < len(newList)-1; i++ {
			if newList[i] > newList[i+1] {
				old := newList[i]
				newList[i] = newList[i+1]
				newList[i+1] = old
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}

	return newList
}

func BubbleByte(u []byte) (newList []byte) {
	newList = make([]byte, len(u))
	copy(newList, u)

	for {
		hasSwap := false
		for i := 0; i < len(newList)-1; i++ {
			if newList[i] > newList[i+1] {
				old := newList[i]
				newList[i] = newList[i+1]
				newList[i+1] = old
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}

	return newList
}

func BubbleRune(u []rune) (newList []rune) {
	newList = make([]rune, len(u))
	copy(newList, u)

	for {
		hasSwap := false
		for i := 0; i < len(newList)-1; i++ {
			if newList[i] > newList[i+1] {
				old := newList[i]
				newList[i] = newList[i+1]
				newList[i+1] = old
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}

	return newList
}

func BubbleUintptr(u []uintptr) (newList []uintptr) {
	newList = make([]uintptr, len(u))
	copy(newList, u)

	for {
		hasSwap := false
		for i := 0; i < len(newList)-1; i++ {
			if newList[i] > newList[i+1] {
				old := newList[i]
				newList[i] = newList[i+1]
				newList[i+1] = old
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}

	return newList
}
