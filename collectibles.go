package collectibles

// Collection is to be exported
type Collection interface {
	Index()
	Include()
	IsConsecutive()
	Compare()
	Map()
	Filter()
	All()
	Any()
}

// Byt type is to be exported
type Byt byte

// Intr type is to be exported
type Intr int

// Flt type is to be exported
type Flt float64

/////////////// byte collection methods Implementation /////////////////

// Index implementation of type byte
func (byt *Byt) Index(vs []byte, t byte) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Map implementation of type byte
func (byt *Byt) Map(vs []byte, f func(byte) byte) []byte {
	vsm := make([]byte, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// Include implementation of type byte
func (byt *Byt) Include(vs []byte, t byte) bool {
	return byt.Index(vs, t) >= 0
}

// Any implementation of type byte
func (byt *Byt) Any(vs []byte, f func(byte) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All implementation of type byte
func (byt *Byt) All(vs []byte, f func(byte) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter implementation of type byte
func (byt *Byt) Filter(vs []byte, f func(byte) bool) []byte {
	vsf := make([]byte, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Compare implementation of type byte
func (byt *Byt) Compare(vs, vp []byte, f func(byte, byte) bool) bool {
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

// IsConsecutive implementation of type int
func (intr *Intr) IsConsecutive(m []int) bool {
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

// Index implementation of type int
func (intr *Intr) Index(vs []int, t int) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Map implementation of type int
func (intr *Intr) Map(vs []int, f func(int) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// Include implementation of type int
func (intr *Intr) Include(vs []int, t int) bool {
	return intr.Index(vs, t) >= 0
}

// Any implementation of type int
func (intr *Intr) Any(vs []int, f func(int) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All implementation of type int
func (intr *Intr) All(vs []int, f func(int) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter implementation of type int
func (intr *Intr) Filter(vs []int, f func(int) bool) []int {
	vsf := make([]int, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Compare implementation of type int
func (intr *Intr) Compare(vs, vp []int, f func(int, int) bool) bool {
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
