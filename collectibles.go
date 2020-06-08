package collectibles

// Collection is to be exported
type Collection interface {
	index()
	include()
	isConsecutive()
	compare()
	mapM()
	filter()
	all()
	any()
}

// Byt type is to be exported
type Byt byte

// Intr type is to be exported
type Intr int

// Flt type is to be exported
type Flt float64

/////////////// byte collection methods Implementation /////////////////

func (byt *Byt) index(vs []byte, t byte) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func (byt *Byt) mapM(vs []byte, f func(byte) byte) []byte {
	vsm := make([]byte, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func (byt *Byt) include(vs []byte, t byte) bool {
	return byt.index(vs, t) >= 0
}

func (byt *Byt) any(vs []byte, f func(byte) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func (byt *Byt) all(vs []byte, f func(byte) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func (byt *Byt) filter(vs []byte, f func(byte) bool) []byte {
	vsf := make([]byte, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func (byt *Byt) compare(vs, vp []byte, f func(byte, byte) bool) bool {
	var retVal bool

	if len(vs) == len(vp) {
		for i := range vs {

			if f(vs[i], vp[i]) {
				retVal = true
			}
			if f(vs[i], vp[i]) != true {
				return false
			}
		}
	}

	if len(vs) != len(vp) {
		return false
	}
	return retVal
}

/////////////// int collection methods Implementation /////////////////

func (intr *Intr) isConsecutive(m []int) bool {
	var retVal bool
	for i, v := range m {
		if i == v {
			retVal = true
		}
		if i != v {
			return false
		}
	}
	return retVal
}

func (intr *Intr) index(vs []int, t int) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func (intr *Intr) mapM(vs []int, f func(int) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func (intr *Intr) include(vs []int, t int) bool {
	return intr.index(vs, t) >= 0
}

func (intr *Intr) any(vs []int, f func(int) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func (intr *Intr) all(vs []int, f func(int) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func (intr *Intr) filter(vs []int, f func(int) bool) []int {
	vsf := make([]int, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func (intr *Intr) compare(vs, vp []int, f func(int, int) bool) bool {
	var retVal bool

	if len(vs) == len(vp) {
		for i := range vs {

			if f(vs[i], vp[i]) {
				retVal = true
			}
			if f(vs[i], vp[i]) != true {
				return false
			}
		}
	}

	if len(vs) != len(vp) {
		return false
	}
	return retVal
}

/////////////// End of int collection methods Implementation ////////////
