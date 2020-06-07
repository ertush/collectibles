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

// Byte_ type is to be exported
type Byte_ byte

// Int_ type is to be exported
type Int_ int

// Float_ type is to be exported
type Float_ float64

/////////////// byte collection methods Implementation /////////////////

func (byte_ *Byte_) index(vs []byte, t byte) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func (byte_ *Byte_) Map(vs []byte, f func(byte) byte) []byte {
	vsm := make([]byte, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func (byte_ *Byte_) include(vs []byte, t byte) bool {
	return byte_.index(vs, t) >= 0
}

func (byte_ *Byte_) any(vs []byte, f func(byte) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func (byte_ *Byte_) all(vs []byte, f func(byte) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func (byte_ *Byte_) filter(vs []byte, f func(byte) bool) []byte {
	vsf := make([]byte, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func (byte_ *Byte_) compare(vs, vp []byte, f func(byte, byte) bool) bool {
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

func (int_ *Int_) isConsecutive(m []int) bool {
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

func (int_ *Int_) index(vs []int, t int) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func (int_ *Int_) Map(vs []int, f func(int) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func (int_ *Int_) include(vs []int, t int) bool {
	return int_.index(vs, t) >= 0
}

func (int_ *Int_) any(vs []int, f func(int) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func (int_ *Int_) all(vs []int, f func(int) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func (int_ *Int_) filter(vs []int, f func(int) bool) []int {
	vsf := make([]int, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func (int_ *Int_) compare(vs, vp []int, f func(int, int) bool) bool {
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
