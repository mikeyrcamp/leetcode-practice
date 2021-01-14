package fizzbuzz

import (
	"strconv"
	"strings"
)

func FizzBuzz(num int) []string {
	result := make([]string, num)
	for i := 1; i <= num; i++ {
		var b strings.Builder
		fizz := i%3 == 0
		buzz := i%5 == 0
		if fizz && buzz {
			b.WriteString("FizzBuzz")
		} else if fizz {
			b.WriteString("Fizz")
		} else if buzz {
			b.WriteString("Buzz")
		} else {
			b.WriteString(strconv.Itoa(i))
		}
		result[i-1] = b.String()
	}
	return result
}
