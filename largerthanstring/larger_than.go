package largerthanstring

// Edge cases: empty strings
// Sample inputs only non-negative numbers
// ["1234", "123"] -> true
// ["123", "1234"] -> false
// ["2345", "2341"]
func LargerThan(a, b string) bool {
	if len(a) == 0 {
		return false
	} else if len(a) != len(b) {
		// a is longer and therefore has a greater value
		return len(a) > len(b)
	}

	// range over the runes in a and b comparing them
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}

	return false
}
