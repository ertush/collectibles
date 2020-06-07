package collectibles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Byte type vars
var byteSlice = []byte{3, 5, 7, 9, 8, 2}
var byteSlice2 = []byte{3, 5, 7, 4, 8, 2}
var byt *Byte_

// Int type vars
var intSlice = []int{3, 5, 7, 9, 8, 2}
var intSlice2 = []int{3, 5, 7, 4, 8, 2}
var intr *Int_

//Byte test help methods

func square(b byte) byte {
	return b * b
}

func byteExist(b byte) bool {
	retVal := false
	if b != 0 {
		return true
	}
	if b == 0 {
		retVal = false
	}
	return retVal
}

func byteExistXclusive(b byte) bool {
	retVal := false
	if b != 0 {
		retVal = true
	}
	if b == 0 {
		return false
	}
	return retVal
}

func isEven(b byte) bool {
	if rem := b % 2; rem == 0 {
		return true
	}
	return false
}

func isEqual(b, c byte) bool {
	if b == c {
		return true
	}
	return false
}

// Int type help methods

func intSquare(b int) int {
	return b * b
}

func intExist(b int) bool {
	retVal := false
	if b != 0 {
		return true
	}
	if b == 0 {
		retVal = false
	}
	return retVal
}

func intExistXclusive(b int) bool {
	retVal := false
	if b != 0 {
		retVal = true
	}
	if b == 0 {
		return false
	}
	return retVal
}

func intIsEven(b int) bool {
	if rem := b % 2; rem == 0 {
		return true
	}
	return false
}

func intIsEqual(b, c int) bool {
	if b == c {
		return true
	}
	return false
}

// Byte type help methods

// Index Test for byte type
func TestIndexForByte(t *testing.T) {
	expect := 2
	assert.Equal(t, expect, byt.index(byteSlice, 7))
}

// Map Test for byte type
func TestMapForByte(t *testing.T) {
	expect := []byte{9, 25, 49, 81, 64, 4}

	assert.Equal(t, expect, byt.Map(byteSlice, square))
}

//Include Test for byte type
func TestIncludeForByte(t *testing.T) {
	expect := true
	assert.Equal(t, expect, byt.include(byteSlice, byte(5)))
}

//Any test for byte type
func TestAnyForByte(t *testing.T) {
	expect := true
	assert.Equal(t, expect, byt.any(byteSlice, byteExist))
}

//All test for byte type
func TestAllForByte(t *testing.T) {
	expect := true
	assert.Equal(t, expect, byt.all(byteSlice, byteExistXclusive))
}

//Filter test for byte type
func TestFilterForByte(t *testing.T) {
	expect := []byte{8, 2}
	assert.Equal(t, expect, byt.filter(byteSlice, isEven))
}

//Compare test for byte type
func TestCompareForByte(t *testing.T) {
	expect := false
	assert.Equal(t, expect, byt.compare(byteSlice, byteSlice2, isEqual))
}

// Int Test Methods

// Index Test for int type
func TestIndexForInt(t *testing.T) {
	expect := 2
	assert.Equal(t, expect, intr.index(intSlice, 7))
}

// Map Test for int type
func TestMapForInt(t *testing.T) {
	expect := []int{9, 25, 49, 81, 64, 4}

	assert.Equal(t, expect, intr.Map(intSlice, intSquare))
}

//Include Test for int type
func TestIncludeForInt(t *testing.T) {
	expect := true
	assert.Equal(t, expect, intr.include(intSlice, 5))
}

//Any test for int type
func TestAnyForInt(t *testing.T) {
	expect := true
	assert.Equal(t, expect, intr.any(intSlice, intExist))
}

//All test for int type
func TestAllForInt(t *testing.T) {
	expect := true
	assert.Equal(t, expect, intr.all(intSlice, intExistXclusive))
}

//Filter test for int type
func TestFilterForInt(t *testing.T) {
	expect := []int{8, 2}
	assert.Equal(t, expect, intr.filter(intSlice, intIsEven))
}

//Compare test for int type
func TestCompareForInt(t *testing.T) {
	expect := false
	assert.Equal(t, expect, intr.compare(intSlice, intSlice2, intIsEqual))
}
