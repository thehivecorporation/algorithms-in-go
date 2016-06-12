package sort

func ShellUint8(u []uint8, j int) (newL []uint8) {
	if j == 0 {
		j = 4
	}

	newL = make([]uint8, len(u))
	copy(newL, u)

	//Interval of 4
	interval := j << 1

	for interval>>1 > 0 {
		interval = interval >> 1

		for i := 0; i < len(newL)-interval; i++ {
			left := newL[i]
			right := newL[i+interval]

			if right < left {
				newL[i] = right
				newL[i+interval] = left
			}
		}
	}

	return
}

func ShellInt(u []int, jump int) (newL []int) {
	if jump == 0 {
		jump = 4
	}

	newL = make([]int, len(u))
	copy(newL, u)

	//Interval of 4
	interval := jump << 1

	for interval>>1 > 0 {
		interval = interval >> 1

		for i := 0; i < len(newL)-interval; i++ {
			left := newL[i]
			right := newL[i+interval]

			if right < left {
				newL[i] = right
				newL[i+interval] = left
			}
		}
	}

	return
}
