package main

func swapInt(u []int, a, b int) {
	old := u[b]
	u[b] = u[a]
	u[a] = old
}

func isLessInt(a, b int) bool {
	if a < b {
		return true
	} else {
		return false
	}
}

var swapTmpl string = `func Swap{{.TypeName}}(u []{{.Type}}, a, b int) {
	old := u[b]
	u[b] = u[a]
	u[a] = old
}

`

var swapTest string = `func TestSwap{{.TypeName}}(t *testing.T) {
	ar := []{{.Type}}{1, 2}

	Swap{{.TypeName}}(ar, 0, 1)
	if ar[0] != 2 || ar[1] != 1 {
		t.Fail()
	}
}

`

var isLessTmpl string = `func IsLess{{.TypeName}}(a, b {{.Type}}) bool {
	if a < b {
		return true
	} else {
		return false
	}
}

`

var isLessTest string = `func TestIsLess{{.TypeName}}(t *testing.T) {
	if !IsLess{{.TypeName}}(1,2) {
		t.Fail()
	}

	if IsLess{{.TypeName}}(2,1) {
		t.Fail()
	}
}

`
