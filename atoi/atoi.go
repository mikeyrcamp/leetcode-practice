package atoi

import (
	"fmt"
	"strings"
)

// Std Cases: ["0", "0.00", "1234", "-1234"]
// Edge Cases: ["", "--1000.00...", "-"]
// If error return 0 with error
func Atoi(a string) (int, error) {
	s := strings.TrimSpace(a)
	s = strings.Split(s, ".")[0]

	sign := 1
	if strings.HasPrefix(s, "-") {
		sign = sign * -1
		s = strings.TrimPrefix(s, "-")
	}
	if s == "" {
		return 0, fmt.Errorf("%s is not a valid integer", a)
	}

	var sum int
	power := 1
	sRunes := []rune(s)
	for i := len(sRunes) - 1; i >= 0; i-- {
		v := int(sRunes[i] - '0')
		if v > 9 || v < 0 {
			return 0, fmt.Errorf("%s is not a valid integer", a)
		}
		// TODO need to handle overflow so this is rather naive
		sum = sum + (power * v)
		power *= 10
	}

	return sum * sign, nil
}
