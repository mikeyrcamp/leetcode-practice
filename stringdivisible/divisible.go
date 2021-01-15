package stringdivisible

import (
	"math"
	"reflect"
	"unsafe"
)

// return -1 if s is divisible by t
// otherwise return the smallest divisor
// see test cases
func divisible(s string, t string) int {
	// First check if t divides s additionally we'll convert these to
	// byte slices
	tBytes := string2bytes(t)
	sBytes := string2bytes(s)
	if !isDivisible(sBytes, tBytes) {
		return -1
	}

	// Is t divisible by a substring of itself that is how
	// we determine what the smallest divisor is
	// we only have to go to the square root of t's length
	maxFactor := int(math.Sqrt(float64(len(t))))
	for i := 1; i <= maxFactor; i++ {
		if isDivisible(tBytes, tBytes[:i]) {
			return i
		}
	}

	return len(t)
}

func isDivisible(s []byte, t []byte) bool {
	if len(s) == 0 && len(t) == 0 || len(s)%len(t) != 0 {
		return false
	}

	// First check if t divides s
	for i := 0; i < len(s); {
		// Compare the bytes from t to s
		for j := 0; j < len(t); i, j = i+1, j+1 {
			if t[j] != s[i] {
				return false
			}
		}
	}
	return true
}

// string2bytes perform a zero copy string to byte slice conversion
// https://www.programmersought.com/article/40465936725/
func string2bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&bh))
}
