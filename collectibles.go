package collectibles

type collection interface {
	index()
	include()
	isConsecutive()
	compare()
	Map()
	filter()
	all()
	any()
}

// _Byte type is to be exported
type _Byte byte

// _Int type is to be exported
type _Int int

// _float type is to be exported
type _float float64

/////////////// byte collection methods Implementation /////////////////

func (byte_ *_Byte) index(vs []byte, t byte) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func (byte_ *_Byte) Map(vs []byte, f func(byte) byte) []byte {
	vsm := make([]byte, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func (byte_ *_Byte) include(vs []byte, t byte) bool {
	return byte_.index(vs, t) >= 0
}

func (byte_ *_Byte) any(vs []byte, f func(byte) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func (byte_ *_Byte) all(vs []byte, f func(byte) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func (byte_ *_Byte) filter(vs []byte, f func(byte) bool) []byte {
	vsf := make([]byte, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func (byte_ *_Byte) compare(vs, vp []byte, f func(byte, byte) bool) bool {
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

func (int_ *_Int) isConsecutive(m []int) bool {
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

func (int_ *_Int) index(vs []int, t int) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func (int_ *_Int) Map(vs []int, f func(int) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func (int_ *_Int) include(vs []int, t int) bool {
	return int_.index(vs, t) >= 0
}

func (int_ *_Int) any(vs []int, f func(int) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func (int_ *_Int) all(vs []int, f func(int) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func (int_ *_Int) filter(vs []int, f func(int) bool) []int {
	vsf := make([]int, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func (int_ *_Int) compare(vs, vp []int, f func(int, int) bool) bool {
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
