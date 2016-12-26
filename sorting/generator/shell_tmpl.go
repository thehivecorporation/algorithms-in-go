package main

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

var shellImpl string = `func Shell{{.TypeName}}(u []{{.Type}}, j int) (newL []{{.Type}}) {
	if j == 0 {
		j = 4
	}

	newL = make([]{{.Type}}, len(u))
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

`

var shellTest string = `func TestShell{{.TypeName}}(t *testing.T) {
	list := []{{.Type}}{4, 2, 1, 3, 0, 9, 6, 8, 7, 5}
	res := Shell{{.TypeName}}(list, 0)

	if len(res) != len(list) {
		t.Fatalf("Returned list has a different number of items %d!=%d",
			len(list), len(res))
	}

	for i := 0; i < len(list); i++ {
		if res[i] != {{.Type}}(i) {
			t.Fatalf("Shell algorithm isn't correct: %v\n", res)
		}
	}
}

`
