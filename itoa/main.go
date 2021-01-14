package itoa

func itoa(v int) string {
	var result []rune
	negative := v < 0
	if negative {
		v = v * -1
	}

	for cur := v; cur != 0; cur = cur / 10 {
		remainder := cur % 10
		result = append([]rune{'0' + rune(remainder)}, result...)
	}

	if negative {
		result = append([]rune{'-'}, result...)
	}

	return string(result)
}
