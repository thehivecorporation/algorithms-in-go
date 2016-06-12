package sort

func BubbleUint8(u []uint8) []uint8 {
	newList := make([]uint8, len(u))
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

func BubbleUint16(u []uint16) []uint16 {
	newList := make([]uint16, len(u))
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

func BubbleUint32(u []uint32) []uint32 {
	newList := make([]uint32, len(u))
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

func BubbleUint64(u []uint64) []uint64 {
	newList := make([]uint64, len(u))
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
