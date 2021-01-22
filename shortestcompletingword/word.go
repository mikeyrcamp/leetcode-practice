package shortestcompletingword

import "unicode"

func shortestCompletingWord(licensePlate string, words []string) string {
	var licOccur [26]int32
	for _, v := range licensePlate {
		if unicode.IsLetter(v) {
			idx := unicode.ToLower(v) - 'a'
			licOccur[idx] = licOccur[idx] + 1
		}
	}

	var shortestMatch string
	for _, w := range words {
		var wOccur [26]int32
		for _, v := range w {
			if unicode.IsLetter(v) {
				idx := unicode.ToLower(v) - 'a'
				wOccur[idx] = wOccur[idx] + 1
			}
		}

		match := true
		for i := range licOccur {
			if licOccur[i] != 0 && wOccur[i] < licOccur[i] {
				match = false
				break
			}
		}

		if match && (len(shortestMatch) > len(w) || shortestMatch == "") {
			shortestMatch = w
		}
	}

	return shortestMatch
}
